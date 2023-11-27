package scylla_test

import (
	"context"
	"testing"

	"github.com/scylladb/gocqlx/v2/gocqlxtest"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla"
)

var userSchemaUP = `
CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY,
  name TEXT,
) WITH compaction = { 'class' : 'LeveledCompactionStrategy' };
`

var userSchemaDown = `
DROP TABLE IF NOT EXISTS gocqlx_test.users;
`

var userCreateStmt = `
INSERT INTO users (
	id,
	name
) VALUES (
	52b23152-0ec1-46d0-b239-b44b392a0485,
	'user'
)
`

func TestUserCreate(t *testing.T) {
	session := gocqlxtest.CreateSession(t)
	session.ExecStmt(userSchemaUP)
	defer session.ExecStmt(userSchemaDown)

	s := scylla.NewUserSession(&session)

	t.Run("save", func(t *testing.T) {
		_, err := s.
			Create().
			SetName("user").
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
