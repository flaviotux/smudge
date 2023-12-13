package scylla

import (
	"context"
	"errors"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"gitlab.luizalabs.com/luizalabs/smudge/graph/model"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/repositories/scylla/user"
)

type UserCreate struct {
	ib       *qb.InsertBuilder
	mutation *UserMutation
	session  *gocqlx.Session
}

func newUserInsertBuilder() *qb.InsertBuilder {
	return qb.Insert(user.Name).Columns(user.Columns...)
}

func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

func (uc *UserCreate) Save(ctx context.Context) (*model.User, error) {
	return uc.cqlSave(ctx)
}

func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`smudge: missing required field "User.name"`)}
	}
	return nil
}

func (uc *UserCreate) cqlSave(ctx context.Context) (*model.User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node := uc.createSpec()
	q := uc.ib.QueryContext(ctx, *uc.session).BindStruct(_node)
	if err := q.ExecRelease(); err != nil {
		return nil, err
	}
	uc.mutation.id = &_node.ID
	return _node, nil
}

func (uc *UserCreate) createSpec() *model.User {
	var _node = &model.User{
		ID: gocql.UUIDFromTime(time.Now()).String(),
	}
	if value, ok := uc.mutation.Name(); ok {
		_node.Name = value
	}
	return _node
}
