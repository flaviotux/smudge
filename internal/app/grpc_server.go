package app

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"gitlab.luizalabs.com/luizalabs/smudge/db"
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
	u := model.User{ID: in.UserID}
	t := model.Todo{
		ID:     gocql.UUIDFromTime(time.Now()).String(),
		Text:   in.Text,
		UserID: u.ID,
	}

	tq := db.TodoTable.InsertQueryContext(ctx, *s.DB).BindStruct(t)
	if err := tq.ExecRelease(); err != nil {
		return nil, err
	}

	uq := db.UserTable.GetQueryContext(ctx, *s.DB).BindStruct(u)
	if err := uq.Get(&u); err != nil {
		return nil, err
	}

	return &proto.TodoResponse{
		ID:   t.ID,
		Text: t.Text,
		Done: t.Done,
		User: &proto.User{
			ID:   u.ID,
			Name: u.Name,
		},
	}, nil
}

func (s *GRPCTodoFetcherServer) GetTodo(ctx context.Context, in *proto.TodoRequest) (*proto.TodosResponse, error) {
	var (
		ts    []model.Todo
		todos = make([]*proto.TodoResponse, 0)
	)

	tq := db.TodoTable.SelectQueryContext(ctx, *s.DB).BindMap(qb.M{})
	if err := tq.SelectRelease(&ts); err != nil {
		return nil, err
	}

	for _, t := range ts {
		u := model.User{ID: t.UserID}
		uq := db.UserTable.GetQueryContext(ctx, *s.DB).BindStruct(u)
		if err := uq.Get(&u); err != nil {
			return nil, err
		}

		todo := &proto.TodoResponse{
			ID:   t.ID,
			Text: t.Text,
			Done: t.Done,
			User: &proto.User{
				ID:   u.ID,
				Name: u.Name,
			},
		}

		todos = append(todos, todo)
	}

	return &proto.TodosResponse{Todos: todos}, nil
}
