package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	"github.com/jice36/blog/models"
	pb "github.com/jice36/blog/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServiceDB struct {
	Conn   *grpc.ClientConn
	client pb.ServiceDBClient
}

func NewService() *ServiceDB {
	conn, err := grpc.Dial("127.0.0.1:5300", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

//	defer conn.Close()

	cli := pb.NewServiceDBClient(conn)

	s := &ServiceDB{
		Conn:   conn,
		client: cli,
	}
	return s
}

func (s *ServiceDB) GetNotes(id string) (*pb.ResponseGet, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()

	r := &pb.RequestGet{Id: id}

	resp, err := s.client.GetNotes(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServiceDB) SendNote(r *models.SendNoteReq) (*pb.ResponseSend, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()

	n := &pb.Note{
		Header:     r.Header,
		Text:       r.Text,
		Tags:       r.Tags,
		TimeCreate: timestamppb.New(time.Now()),
	}

	rs := &pb.RequestSend{
		Id: r.UserID,
		Note: n,
	}

	resp, err := s.client.SendNote(ctx, rs)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return resp, nil
}
