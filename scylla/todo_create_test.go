package scylla_test

import (
	"context"
	"testing"

	"github.com/scylladb/gocqlx/v2/gocqlxtest"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla"
)

func TestTodoCreate(t *testing.T) {
	session := gocqlxtest.CreateSession(t)
	session.ExecStmt(todoSchemaUP)
	defer session.ExecStmt(todoSchemaDown)

	s := scylla.NewTodoSession(&session)

	t.Run("simple", func(t *testing.T) {
		_, err := s.
			Create().
			SetText("text").
			SetDone(false).
			SetUserId("52b23152-0ec1-46d0-b239-b44b392a0485").
			Save(context.Background())
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("check required fields", func(t *testing.T) {
		_, err := s.
			Create().
			Save(context.Background())
		if err == nil {
			t.Fatal(err)
		}
	})
}
