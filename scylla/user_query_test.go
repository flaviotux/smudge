package scylla_test

import (
	"context"
	"testing"

	"github.com/scylladb/gocqlx/v2/gocqlxtest"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla/user"
)

func TestUserQuery(t *testing.T) {
	session := gocqlxtest.CreateSession(t)
	session.ExecStmt(userSchemaUP)
	defer session.ExecStmt(userSchemaDown)

	s := scylla.NewUserSession(&session)

	t.Run("save", func(t *testing.T) {
		if err := session.ExecStmt(userCreateStmt); err != nil {
			t.Fatal(err)
		}

		user, err := s.
			Query().
			Where(user.ID("52b23152-0ec1-46d0-b239-b44b392a0485")).
			Only(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != "52b23152-0ec1-46d0-b239-b44b392a0485" {
			t.Fail()
		}
	})

	t.Run("all", func(t *testing.T) {
		if err := session.ExecStmt(userCreateStmt); err != nil {
			t.Fatal(err)
		}

		user, err := s.
			Query().
			Where(user.ID("52b23152-0ec1-46d0-b239-b44b392a0485")).
			All(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if user[0].ID != "52b23152-0ec1-46d0-b239-b44b392a0485" {
			t.Fail()
		}
	})
}
