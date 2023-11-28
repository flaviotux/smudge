package scylla

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"gitlab.luizalabs.com/luizalabs/smudge/graph/model"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla/todo"
)

type TodoUpdate struct {
	ub       *qb.UpdateBuilder
	mutation *TodoMutation
	session  *gocqlx.Session
}

func newTodoUpdateBuilder() *qb.UpdateBuilder {
	return qb.Update(todo.Table)
}

// SetText sets the "text" field.
func (tu *TodoUpdate) SetText(s string) *TodoUpdate {
	tu.mutation.SetText(s)
	return tu
}

// SetDone sets the "done" field.
func (tu *TodoUpdate) SetDone(b bool) *TodoUpdate {
	tu.mutation.SetDone(b)
	return tu
}

// SetUserId sets the "user_id" field.
func (tu *TodoUpdate) SetUserId(s string) *TodoUpdate {
	tu.mutation.SetUserId(s)
	return tu
}

// AddUser adds the "user" children to the Todo entity.
func (tu *TodoUpdate) AddUser(u *model.User) *TodoUpdate {
	return tu.SetUserId(u.ID)
}

// Where appends a list comparators to the TodoUpdate builder.
func (tu *TodoUpdate) Where(ps ...qb.Cmp) *TodoUpdate {
	tu.ub.Where(ps...)
	return tu
}

// Mutation returns the TodoMutation object of the builder.
func (tu *TodoUpdate) Mutation() *TodoMutation {
	return tu.mutation
}

func (tu *TodoUpdate) Save(ctx context.Context) (int, error) {
	return tu.cqlSave(ctx)
}

func (tu *TodoUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

func (tu *TodoUpdate) cqlSave(ctx context.Context) (int, error) {
	var (
		_node   = tu.createSpec()
		q       = tu.ub.QueryContext(ctx, *tu.session).BindStruct(_node)
		counter = q.Iter().NumRows()
	)
	if err := q.ExecRelease(); err != nil {
		return 0, err
	}
	return counter, nil
}

func (tu *TodoUpdate) createSpec() *model.Todo {
	_node := &model.Todo{}
	if value, ok := tu.mutation.ID(); ok {
		_node.ID = value
	}
	if value, ok := tu.mutation.Text(); ok {
		tu.ub.Set(todo.FieldText)
		_node.Text = value
	}
	if value, ok := tu.mutation.Done(); ok {
		tu.ub.Set(todo.FieldDone)
		_node.Done = value
	}
	if value, ok := tu.mutation.UserID(); ok {
		tu.ub.Set(todo.FieldUserID)
		_node.UserID = value
	}

	return _node
}

type TodoUpdateOne struct {
	tu *TodoUpdate
}

// SetText sets the "text" field.
func (tuo *TodoUpdateOne) SetText(s string) *TodoUpdateOne {
	tuo.tu.mutation.SetText(s)
	return tuo
}

// SetDone sets the "done" field.
func (tuo *TodoUpdateOne) SetDone(b bool) *TodoUpdateOne {
	tuo.tu.mutation.SetDone(b)
	return tuo
}

// SetUserId sets the "user_id" field.
func (tuo *TodoUpdateOne) SetUserId(s string) *TodoUpdateOne {
	tuo.tu.mutation.SetUserId(s)
	return tuo
}

// AddUser adds the "user" children to the Todo entity.
func (tuo *TodoUpdateOne) AddUser(u *model.User) *TodoUpdateOne {
	return tuo.SetUserId(u.ID)
}

// Where appends a list predicates to the TodoUpdate builder.
func (tuo *TodoUpdateOne) Where(ps ...qb.Cmp) *TodoUpdateOne {
	tuo.tu.Where(ps...)
	return tuo
}

func (tuo *TodoUpdateOne) Save(ctx context.Context) error {
	n, err := tuo.tu.Save(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{todo.Label, *tuo.tu.mutation.id}
	default:
		return nil
	}
}
