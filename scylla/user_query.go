package scylla

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"gitlab.luizalabs.com/luizalabs/smudge/graph/model"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla/user"
)

type UserQuery struct {
	sb      *qb.SelectBuilder
	session *gocqlx.Session
}

func newUserSelectBuilder() *qb.SelectBuilder {
	return qb.Select(user.Table)
}

func (uq *UserQuery) Columns(columns ...string) *UserQuery {
	uq.sb.Columns(columns...)
	return uq
}

func (uq *UserQuery) Where(w ...qb.Cmp) *UserQuery {
	uq.sb.Where(w...)
	return uq
}

func (uq *UserQuery) Limit(limit uint) *UserQuery {
	uq.sb.Limit(limit)
	return uq
}

func (uq *UserQuery) OrderBy(column string, o qb.Order) *UserQuery {
	uq.sb.OrderBy(column, o)
	return uq
}

func (uq *UserQuery) Only(ctx context.Context) (*model.User, error) {
	var user model.User
	u := uq.sb.QueryContext(ctx, *uq.session).BindStruct(user)
	if err := u.GetRelease(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (uq *UserQuery) All(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	q := uq.sb.QueryContext(ctx, *uq.session)
	if err := q.SelectRelease(&users); err != nil {
		return nil, err
	}
	return users, nil
}
