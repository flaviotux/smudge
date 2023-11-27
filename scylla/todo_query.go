package scylla

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"gitlab.luizalabs.com/luizalabs/smudge/graph/model"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla/todo"
)

type TodoQuery struct {
	sb      *qb.SelectBuilder
	session *gocqlx.Session
}

func newTodoSelectBuilder() *qb.SelectBuilder {
	return qb.Select(todo.Table)
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
	var todo model.Todo
	u := tq.sb.QueryContext(ctx, *tq.session).BindStruct(todo)
	if err := u.GetRelease(&todo); err != nil {
		return nil, err
	}
	return &todo, nil
}

func (tq *TodoQuery) All(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	q := tq.sb.QueryContext(ctx, *tq.session)
	if err := q.SelectRelease(&todos); err != nil {
		return nil, err
	}
	return todos, nil
}
