package db

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/model"
)

type UserQuery struct {
	table.Table

	session *gocqlx.Session
	where   []qb.Cmp
	limit   uint
}

func (uq *UserQuery) Where(w ...qb.Cmp) *UserQuery {
	uq.where = append(uq.where, w...)
	return uq
}

func (uq *UserQuery) Limit(limit uint) *UserQuery {
	uq.limit = limit
	return uq
}

func (uq *UserQuery) Only(ctx context.Context) (*model.User, error) {
	var user model.User
	sb := uq.SelectBuilder()
	if uq.where != nil {
		sb.Where(uq.where...)
	}
	if uq.limit != 0 {
		sb.Limit(uq.limit)
	}
	u := sb.QueryContext(ctx, *uq.session).BindStruct(user)
	if err := u.GetRelease(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (uq *UserQuery) All(ctx context.Context) ([]*model.User, error) {
	var user []*model.User
	sb := uq.SelectBuilder()
	if uq.where != nil {
		sb.Where(uq.where...)
	}
	if uq.limit != 0 {
		sb.Limit(uq.limit)
	}
	q := sb.QueryContext(ctx, *uq.session)
	if err := q.SelectRelease(&user); err != nil {
		return nil, err
	}
	return user, nil
}
