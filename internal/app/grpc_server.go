package app

import (
	"context"
	"log"
	"net"

	"gitlab.luizalabs.com/luizalabs/smudge/db"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/model"
	"gitlab.luizalabs.com/luizalabs/smudge/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultGRPCAddr = ":8090"

func MakeGRPCServerAndRun(listenAddr string, session *db.Session) error {
	if listenAddr == "" {
		listenAddr = defaultGRPCAddr
	}

	grpcTodoFetcher := NewGRPCTodoFetcherServer(session)

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
	session *db.Session
}

func NewGRPCTodoFetcherServer(session *db.Session) *GRPCTodoFetcherServer {
	return &GRPCTodoFetcherServer{session: session}
}

func (s *GRPCTodoFetcherServer) NewTodo(ctx context.Context, in *proto.TodoRequest) (*proto.TodoResponse, error) {
	user := model.User{ID: in.UserID}

	if _, err := s.session.User.Query().Only(ctx); err != nil {
		return nil, err
	}

	t, err := s.session.Todo.
		Create().
		SetText(in.Text).
		SetDone(false).
		AddUser(&user).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return &proto.TodoResponse{
		ID:   t.ID,
		Text: t.Text,
		Done: t.Done,
		User: &proto.User{
			ID:   t.User.ID,
			Name: t.User.Name,
		},
	}, nil
}

func (s *GRPCTodoFetcherServer) GetTodo(ctx context.Context, in *proto.TodoRequest) (*proto.TodosResponse, error) {
	var tr = make([]*proto.TodoResponse, 0)
	todos, err := s.session.Todo.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	for _, t := range todos {
		user, err := s.session.User.Get(ctx, t.UserID)
		if err != nil {
			return nil, err
		}

		todo := &proto.TodoResponse{
			ID:   t.ID,
			Text: t.Text,
			Done: t.Done,
			User: &proto.User{
				ID:   user.ID,
				Name: user.Name,
			},
		}

		tr = append(tr, todo)
	}
	return &proto.TodosResponse{Todos: tr}, nil
}
