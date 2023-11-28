package scylla_test

import (
	"context"
	"testing"

	"github.com/scylladb/gocqlx/v2/gocqlxtest"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla/user"
)

func TestUserUpdate(t *testing.T) {
	session := gocqlxtest.CreateSession(t)
	session.ExecStmt(userSchemaUP)
	defer session.ExecStmt(userSchemaDown)

	s := scylla.NewUserSession(&session)

	t.Run("simple", func(t *testing.T) {
		if err := session.ExecStmt(userCreateStmt); err != nil {
			t.Fatal(err)
		}

		i, err := s.
			Update().
			Where(user.ID("52b23152-0ec1-46d0-b239-b44b392a0485")).
			SetName("user 2").
			Save(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if i < 0 {
			t.Fail()
		}
	})

	t.Run("not exists", func(t *testing.T) {
		if err := session.ExecStmt(userCreateStmt); err != nil {
			t.Fatal(err)
		}

		err := s.
			UpdateOneID("d763fe8a-6b5e-414c-a109-3b277f1d0a54").
			SetName("user 2").
			Save(context.Background())
		if err == nil {
			t.Fatal(err)
		}
	})
}
