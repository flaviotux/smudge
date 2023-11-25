package scylla

import (
	"context"
	"errors"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
	"gitlab.luizalabs.com/luizalabs/smudge/graph/model"
)

type TodoCreate struct {
	table.Table

	mutation *TodoMutation
	session  *gocqlx.Session
}

func (tc *TodoCreate) SetText(s string) *TodoCreate {
	tc.mutation.SetText(s)
	return tc
}

func (tc *TodoCreate) SetDone(b bool) *TodoCreate {
	tc.mutation.SetDone(b)
	return tc
}

func (tc *TodoCreate) SetUserId(s string) *TodoCreate {
	tc.mutation.SetUserId(s)
	return tc
}

func (tc *TodoCreate) AddUser(u *model.User) *TodoCreate {
	return tc.SetUserId(u.ID)
}

func (tc *TodoCreate) defaults() {
	tc.mutation.SetDone(false)
}

func (tc *TodoCreate) Save(ctx context.Context) (*model.Todo, error) {
	tc.defaults()
	return tc.cqlSave(ctx)
}

// check runs all checks and user-defined validators on the builder.
func (tc *TodoCreate) check() error {
	if _, ok := tc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`smudge: missing required field "Todo.text"`)}
	}
	if _, ok := tc.mutation.Done(); !ok {
		return &ValidationError{Name: "done", err: errors.New(`smudge: missing required field "Todo.done"`)}
	}
	if _, ok := tc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`smudge: missing required field "Todo.user_id"`)}
	}
	return nil
}

func (tc *TodoCreate) cqlSave(ctx context.Context) (*model.Todo, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node := tc.createSpec()
	q := tc.InsertQueryContext(ctx, *tc.session).BindStruct(_node)
	if err := q.ExecRelease(); err != nil {
		return nil, err
	}
	tc.mutation.id = &_node.ID
	return _node, nil
}

func (tc *TodoCreate) createSpec() *model.Todo {
	var _node = &model.Todo{
		ID: gocql.UUIDFromTime(time.Now()).String(),
	}
	if value, ok := tc.mutation.Text(); ok {
		_node.Text = value
	}
	if value, ok := tc.mutation.Done(); ok {
		_node.Done = value
	}
	if value, ok := tc.mutation.UserID(); ok {
		_node.UserID = value
	}
	return _node
}
