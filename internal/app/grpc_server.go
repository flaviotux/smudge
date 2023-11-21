package app

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"

	"gitlab.luizalabs.com/luizalabs/smudge/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultGRPCAddr = ":8090"

func MakeGRPCServerAndRun(listenAddr string) error {
	if listenAddr == "" {
		listenAddr = defaultGRPCAddr
	}

	grpcTodoFetcher := NewGRPCTodoFetcherServer()

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	reflection.Register(server)
	proto.RegisterTodoServer(server, grpcTodoFetcher)

	log.Printf("Serving gRPC API on 0.0.0.0%s\n", listenAddr)
	return server.Serve(ln)
}

type GRPCTodoFetcherServer struct {
	proto.UnimplementedTodoServer
	todos []*proto.TodoResponse
}

func NewGRPCTodoFetcherServer() *GRPCTodoFetcherServer {
	return &GRPCTodoFetcherServer{}
}

func (s *GRPCTodoFetcherServer) NewTodo(ctx context.Context, in *proto.TodoRequest) (*proto.TodoResponse, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &proto.TodoResponse{
		Text:   in.Text,
		ID:     fmt.Sprintf("T%d", randNumber),
		UserID: in.UserID,
	}

	s.todos = append(s.todos, todo)

	return todo, nil
}

func (s *GRPCTodoFetcherServer) GetTodo(ctx context.Context, in *proto.TodoRequest) (*proto.TodosResponse, error) {
	return &proto.TodosResponse{
		Todos: s.todos,
	}, nil
}
