package scylla_test

import (
	"context"
	"testing"

	"github.com/scylladb/gocqlx/v2/gocqlxtest"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/repositories/scylla"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/repositories/scylla/todo"
)

func TestTodoUpdate(t *testing.T) {
	session := gocqlxtest.CreateSession(t)
	session.ExecStmt(todoSchemaUP)
	defer session.ExecStmt(todoSchemaDown)

	s := scylla.NewTodoSession(&session)

	t.Run("simple", func(t *testing.T) {
		if err := session.ExecStmt(todoCreateStmt); err != nil {
			t.Fatal(err)
		}

		i, err := s.
			Update().
			Where(
				todo.IDEQ("d763fe8a-6b5e-414c-a109-3b277f1d0a54"),
				todo.UserIDEQ("52b23152-0ec1-46d0-b239-b44b392a0485"),
			).
			SetText("text 2").
			Save(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if i < 0 {
			t.Fail()
		}
	})

	t.Run("not exists", func(t *testing.T) {
		if err := session.ExecStmt(todoCreateStmt); err != nil {
			t.Fatal(err)
		}

		err := s.
			UpdateOneID("52b23152-0ec1-46d0-b239-b44b392a0485").
			SetText("user 2").
			Save(context.Background())
		if err == nil {
			t.Fatal(err)
		}
	})
}
