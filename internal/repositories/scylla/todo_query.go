package scylla

import (
	"context"
	"fmt"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"gitlab.luizalabs.com/luizalabs/smudge/graph/model"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/repositories/scylla/todo"
)

type TodoQuery struct {
	sb       *qb.SelectBuilder
	mutation *TodoMutation
	session  *gocqlx.Session
}

func newTodoSelectBuilder() *qb.SelectBuilder {
	return todo.Table.SelectBuilder()
}

func (tq *TodoQuery) Columns(columns ...string) *TodoQuery {
	tq.sb.Columns(columns...)
	return tq
}

func (uq *TodoQuery) Where(w ...qb.Cmp) *TodoQuery {
	uq.sb.Where(w...)
	return uq
}

func (uq *TodoQuery) Limit(limit uint) *TodoQuery {
	uq.sb.Limit(limit)
	return uq
}

func (uq *TodoQuery) OrderBy(column string, o qb.Order) *TodoQuery {
	uq.sb.OrderBy(column, o)
	return uq
}

func (tq *TodoQuery) Only(ctx context.Context) (*model.Todo, error) {
	var t model.Todo
	u := tq.sb.QueryContext(ctx, *tq.session).BindMap(qb.M{
		todo.FieldID:     *tq.mutation.id,
		todo.FieldText:   *tq.mutation.text,
		todo.FieldDone:   *tq.mutation.done,
		todo.FieldUserID: *tq.mutation.user_id,
	})
	fmt.Printf("u.Values(): %v\n", u.Values())
	if err := u.GetRelease(&t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (tq *TodoQuery) All(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo

	q := tq.sb.QueryContext(ctx, *tq.session).BindMap(qb.M{
		todo.FieldID:     tq.mutation.id,
		todo.FieldUserID: tq.mutation.user_id,
		todo.FieldText:   tq.mutation.text,
		todo.FieldDone:   tq.mutation.done,
	})
	if err := q.SelectRelease(&todos); err != nil {
		return nil, err
	}
	return todos, nil
}
