package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/lib/pq"
	"log"
	"net"
	"time"

	c "github.com/jice36/blog/internal/server"
	pb "github.com/jice36/blog/proto"

	"github.com/google/uuid"
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
	flag.StringVar(&configPath, "config", "config/config.yml", "path to config file") //todo пароль поправить в конфиге
}

func main() {
	flag.Parse()

	config, cErr := c.NewConfigDB(configPath)
	if cErr != nil {
		log.Fatal(cErr)
	}

	connStr := "user=" + config.Database.Dbuser + " password=" + config.Database.Dbpassword + " dbname=" + config.Database.Dbname + " sslmode=disable"
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
	if r.Id == "" {
		return nil, errors.New("")
	}
	q := `select header, text, tags, time_create from notes where id_user = $1`

	rows, err := s.DB.Query(q, r.Id) // проверить на пользователя
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

	param := []interface{}{idNote, r.Note.Header, r.Note.Text, pq.Array(r.Note.Tags), time.Now(), r.Id}

	_, err = tx.Exec(q, param...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	a := &pb.ResponseSend{}
	return a, nil
}
