package scylla_test

import (
	"context"
	"testing"

	"github.com/scylladb/gocqlx/v2/gocqlxtest"
	"gitlab.luizalabs.com/luizalabs/smudge/graph/model"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/repositories/scylla"
)

func TestTodoQuery(t *testing.T) {
	session := gocqlxtest.CreateSession(t)
	session.ExecStmt(todoSchemaUP)
	defer session.ExecStmt(todoSchemaDown)

	s := scylla.NewTodoSession(&session)

	t.Run("only", func(t *testing.T) {
		if err := session.ExecStmt(todoCreateStmt); err != nil {
			t.Fatal(err)
		}

		todo, err := s.
			Query(&model.Todo{ID: "d763fe8a-6b5e-414c-a109-3b277f1d0a54", UserID: "52b23152-0ec1-46d0-b239-b44b392a0485", Done: false}).
			Only(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if todo.ID != "d763fe8a-6b5e-414c-a109-3b277f1d0a54" {
			t.Fail()
		}
	})

	t.Run("all", func(t *testing.T) {
		if err := session.ExecStmt(todoCreateStmt); err != nil {
			t.Fatal(err)
		}

		todo, err := s.
			Query(&model.Todo{ID: "d763fe8a-6b5e-414c-a109-3b277f1d0a54", UserID: "52b23152-0ec1-46d0-b239-b44b392a0485"}).
			All(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if todo[0].ID != "d763fe8a-6b5e-414c-a109-3b277f1d0a54" {
			t.Fail()
		}
	})
}
