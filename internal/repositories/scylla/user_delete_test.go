package scylla_test

import (
	"context"
	"testing"

	"github.com/scylladb/gocqlx/v2/gocqlxtest"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/repositories/scylla"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/repositories/scylla/user"
)

func TestUserDelete(t *testing.T) {
	session := gocqlxtest.CreateSession(t)
	session.ExecStmt(userSchemaUP)
	defer session.ExecStmt(userSchemaDown)

	s := scylla.NewUserSession(&session)

	t.Run("simple", func(t *testing.T) {
		if err := session.ExecStmt(userCreateStmt); err != nil {
			t.Fatal(err)
		}

		i, err := s.
			Delete().
			Where(user.ID("d763fe8a-6b5e-414c-a109-3b277f1d0a54")).
			Exec(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if i > 1 {
			t.Fail()
		}
	})

	t.Run("not exists", func(t *testing.T) {
		if err := session.ExecStmt(userCreateStmt); err != nil {
			t.Fatal(err)
		}

		err := s.
			DeleteOneID("52b23152-0ec1-46d0-b239-b44b392a0485").
			Exec(context.Background())
		if err == nil {
			t.Fatal(err)
		}
	})
}
