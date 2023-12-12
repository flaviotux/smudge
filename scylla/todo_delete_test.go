package scylla_test

import (
	"context"
	"testing"

	"github.com/scylladb/gocqlx/v2/gocqlxtest"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla/todo"
)

var todoSchemaUP = `
CREATE TABLE IF NOT EXISTS gocqlx_test.todos (
  id UUID,
  text TEXT,
  done BOOLEAN,
  user_id UUID,
	PRIMARY KEY (id, user_id)
) WITH compaction = { 'class' : 'LeveledCompactionStrategy' };
`

var todoSchemaDown = `
DROP TABLE IF EXISTS gocqlx_test.todos;
`

var todoCreateStmt = `
INSERT INTO todos (
	id,
	text,
	done,
	user_id
) VALUES (
	d763fe8a-6b5e-414c-a109-3b277f1d0a54,
	'text',
	false,
	52b23152-0ec1-46d0-b239-b44b392a0485
)
`

func TestTodoDelete(t *testing.T) {
	session := gocqlxtest.CreateSession(t)
	session.ExecStmt(todoSchemaUP)
	defer session.ExecStmt(todoSchemaDown)

	s := scylla.NewTodoSession(&session)

	t.Run("simple", func(t *testing.T) {
		if err := session.ExecStmt(todoCreateStmt); err != nil {
			t.Fatal(err)
		}

		i, err := s.
			Delete().
			Where(todo.ID("d763fe8a-6b5e-414c-a109-3b277f1d0a54")).
			Exec(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if i > 1 {
			t.Fail()
		}
	})

	t.Run("not exists", func(t *testing.T) {
		if err := session.ExecStmt(todoCreateStmt); err != nil {
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
