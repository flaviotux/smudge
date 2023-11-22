package app

import (
	"context"
	"log"
	"net"

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
	session *gocqlx.Session
}

func NewGRPCTodoFetcherServer(session *gocqlx.Session) *GRPCTodoFetcherServer {
	return &GRPCTodoFetcherServer{session: session}
}

func (s *GRPCTodoFetcherServer) NewTodo(ctx context.Context, in *proto.TodoRequest) (*proto.TodoResponse, error) {
	user := model.User{ID: in.UserID}

	t, err := model.NewTodoModel(s.session).
		SetText(in.Text).
		SetDone(false).
		AddUser(&user).
		InsertQueryContext(ctx)

	if err != nil {
		return nil, err
	}

	u := model.NewUserModel(s.session)
	if _, err := u.GetQueryContext(ctx); err != nil {
		return nil, err
	}

	t.AddUser(u)

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
		ts    = model.NewTodoModel(s.session)
		todos = make([]*proto.TodoResponse, 0)
	)

	t, err := ts.SelectQueryContext(ctx, qb.M{})
	if err != nil {
		return nil, err
	}

	for _, t := range t {
		user := model.NewUserModel(s.session)
		user.ID = t.UserID

		if _, err := user.GetQueryContext(ctx); err != nil {
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
