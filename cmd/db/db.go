package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	c "github.com/jice36/blog/internal/server"
	pb "github.com/jice36/blog/proto"

	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedServiceDBServer
	DB *sql.DB
}

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "config/config.yml", "path to config file")
}

func main() {
	flag.Parse()

	config, cErr := c.NewConfigDB(configPath)
	if cErr != nil {
		log.Fatal(cErr)
	}

	connStr := "host=" + config.Database.Dbhost + " port=" + config.Database.Dbport + " user=" + config.Database.Dbuser +
		" password=" + config.Database.Dbpassword + " dbname=" + config.Database.Dbname + " sslmode=disable"
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterServiceDBServer(grpcServer, &server{
		UnimplementedServiceDBServer: pb.UnimplementedServiceDBServer{},
		DB:                           db,
	})
	grpcServer.Serve(listener)
}

func (s *server) GetNotes(ctx context.Context, r *pb.RequestGet) (*pb.ResponseGet, error) {
	ctxLocal, cancel := context.WithCancel(ctx)
	defer cancel()

	if r.Id == "" {
		return nil, errors.New("id не передан")
	}
	notesResp, err := s.getNotes(ctxLocal, r)
	if err != nil {
		return nil, err
	}

	return notesResp, nil
}

func (s *server) getNotes(ctx context.Context, r *pb.RequestGet) (*pb.ResponseGet, error) {
	q := `select header, text, tags, time_create from notes where id_user = $1`

	rows, err := s.DB.Query(q, r.Id)
	if err != nil {
		return nil, err
	}

	notes := make([]*pb.Note, 0)

	for rows.Next() {
		var header, text sql.NullString
		var tagsS []string
		var timeCreate sql.NullTime

		err := rows.Scan(&header, &text, pq.Array(&tagsS), &timeCreate)
		if err != nil {
			return nil, err
		}

		if !header.Valid || !text.Valid || !timeCreate.Valid {
			return nil, errors.New("не удалось получить данные")
		}

		note := &pb.Note{
			Header:     header.String,
			Text:       text.String,
			Tags:       tagsS,
			TimeCreate: timestamppb.New(timeCreate.Time),
		}
		notes = append(notes, note)
	}

	if len(notes) == 0 {
		return nil, errors.New("записей нет")
	}

	notesResp := &pb.ResponseGet{
		Notes: notes,
	}

	return notesResp, nil
}

func (s *server) SendNote(ctx context.Context, r *pb.RequestSend) (*pb.ResponseSend, error) {
	ctxLocal, cancel := context.WithCancel(ctx)
	defer cancel()

	sendResp, err := s.sendNotes(ctxLocal, r)
	if err != nil {
		return nil, err
	}

	return sendResp, nil
}

func (s *server) sendNotes(ctx context.Context, r *pb.RequestSend) (*pb.ResponseSend, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}

	q := `INSERT INTO notes (id_note, header, text, tags, time_create, id_user) 
				VALUES ($1, $2, $3, $4, $5, $6)`
	idNote, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	param := []interface{}{idNote, r.Note.Header, r.Note.Text, pq.Array(r.Note.Tags), r.Note.TimeCreate.AsTime(), r.Id}

	_, err = tx.Exec(q, param...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	res := &pb.ResponseSend{Done: true}
	return res, nil
}
