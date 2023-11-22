package app

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/model"
	"gitlab.luizalabs.com/luizalabs/smudge/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultGRPCAddr = ":8090"

func MakeGRPCServerAndRun(listenAddr string, session *gocqlx.Session) error {
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
	DB *gocqlx.Session
}

func NewGRPCTodoFetcherServer(session *gocqlx.Session) *GRPCTodoFetcherServer {
	return &GRPCTodoFetcherServer{DB: session}
}

func (s *GRPCTodoFetcherServer) NewTodo(ctx context.Context, in *proto.TodoRequest) (*proto.TodoResponse, error) {
	t, err := model.NewTodoModel(s.DB).
		SetID(gocql.UUIDFromTime(time.Now()).String()).
		SetText(in.Text).
		SetDone(false).
		SetUserID(in.UserID).
		GetQueryContext(ctx)

	if err != nil {
		return nil, err
	}

	if _, err := t.WithUserContext(ctx); err != nil {
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
	var (
		ts    = model.NewTodoModel(s.DB)
		todos = make([]*proto.TodoResponse, 0)
	)

	t, err := ts.SelectQueryContext(ctx, qb.M{})
	if err != nil {
		return nil, err
	}

	for _, t := range t {
		_, err := t.WithUserContext(ctx)
		if err != nil {
			return nil, err
		}

		todo := &proto.TodoResponse{
			ID:   t.ID,
			Text: t.Text,
			Done: t.Done,
			User: &proto.User{
				ID:   t.User.ID,
				Name: t.User.Name,
			},
		}

		todos = append(todos, todo)
	}

	return &proto.TodosResponse{Todos: todos}, nil
}
