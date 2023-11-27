package scylla

import (
	"context"
	"errors"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"gitlab.luizalabs.com/luizalabs/smudge/graph/model"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla/todo"
)

type TodoCreate struct {
	ib       *qb.InsertBuilder
	mutation *TodoMutation
	session  *gocqlx.Session
}

func newTodoInsertBuilder() *qb.InsertBuilder {
	return qb.Insert(todo.Table).Columns(todo.Columns...)
}

// SetText sets the "text" field.
func (tc *TodoCreate) SetText(s string) *TodoCreate {
	tc.mutation.SetText(s)
	return tc
}

// SetDone sets the "done" field.
func (tc *TodoCreate) SetDone(b bool) *TodoCreate {
	tc.mutation.SetDone(b)
	return tc
}

// SetUserId sets the "user_id" field.
func (tc *TodoCreate) SetUserId(s string) *TodoCreate {
	tc.mutation.SetUserId(s)
	return tc
}

// AddUser adds the "user" children to the Todo entity.
func (tc *TodoCreate) AddUser(u *model.User) *TodoCreate {
	return tc.SetUserId(u.ID)
}

// Mutation returns the TodoMutation object of the builder.
func (tc *TodoCreate) Mutation() *TodoMutation {
	return tc.mutation
}

// Save creates the Todo in the database.
func (tc *TodoCreate) Save(ctx context.Context) (*model.Todo, error) {
	tc.defaults()
	return tc.cqlSave(ctx)
}

// defaults sets the default values of the builder before save.
func (tc *TodoCreate) defaults() {
	if _, ok := tc.mutation.Done(); !ok {
		v := todo.DefaultDone
		tc.mutation.SetDone(v)
	}
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
	q := tc.ib.QueryContext(ctx, *tc.session).BindStruct(_node)
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
