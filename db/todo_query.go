package db

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/model"
)

type TodoQuery struct {
	table.Table

	session *gocqlx.Session
	where   []qb.Cmp
	limit   uint
}

func (tq *TodoQuery) Where(w ...qb.Cmp) *TodoQuery {
	tq.where = append(tq.where, w...)
	return tq
}

func (tq *TodoQuery) Limit(i uint) *TodoQuery {
	tq.limit = i
	return tq
}

func (tq *TodoQuery) All(ctx context.Context) ([]*model.Todo, error) {
	var todo []*model.Todo
	sb := tq.SelectBuilder()
	if tq.where != nil {
		sb.Where(tq.where...)
	}
	if tq.limit != 0 {
		sb.Limit(tq.limit)
	}
	q := sb.QueryContext(ctx, *tq.session)
	if err := q.SelectRelease(&todo); err != nil {
		return nil, err
	}
	return todo, nil
}
