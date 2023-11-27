package scylla_test

import (
	"context"
	"testing"

	"github.com/scylladb/gocqlx/v2/gocqlxtest"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla/todo"
)

func TestTodoQuery(t *testing.T) {
	session := gocqlxtest.CreateSession(t)
	session.ExecStmt(todoSchemaUP)
	defer session.ExecStmt(todoSchemaDown)

	t.Run("only", func(t *testing.T) {
		if err := session.ExecStmt(todoCreateStmt); err != nil {
			t.Fatal(err)
		}

		s := scylla.NewTodoSession(&session)
		todo, err := s.
			Query().
			Where(todo.ID("d763fe8a-6b5e-414c-a109-3b277f1d0a54")).
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

		s := scylla.NewTodoSession(&session)
		todo, err := s.
			Query().
			Where(todo.ID("d763fe8a-6b5e-414c-a109-3b277f1d0a54")).
			All(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if todo[0].ID != "d763fe8a-6b5e-414c-a109-3b277f1d0a54" {
			t.Fail()
		}
	})
}
