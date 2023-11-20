package smudge

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"gitlab.luizalabs.com/luizalabs/smudge/pb"
)

type Server struct {
	pb.UnimplementedTodoServer
	todos []*pb.TodoResponse
}

func (s *Server) NewTodo(ctx context.Context, in *pb.TodoRequest) (*pb.TodoResponse, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &pb.TodoResponse{
		Text:   in.Text,
		ID:     fmt.Sprintf("T%d", randNumber),
		UserID: in.UserID,
	}

	s.todos = append(s.todos, todo)

	return todo, nil
}

func (s *Server) GetTodo(ctx context.Context, in *pb.TodoRequest) (*pb.TodosResponse, error) {
	return &pb.TodosResponse{
		Todos: s.todos,
	}, nil
}
