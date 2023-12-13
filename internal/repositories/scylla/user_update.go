package scylla

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"gitlab.luizalabs.com/luizalabs/smudge/graph/model"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/repositories/scylla/user"
)

type UserUpdate struct {
	ub       *qb.UpdateBuilder
	mutation *UserMutation
	session  *gocqlx.Session
}

func newUserUpdateBuilder() *qb.UpdateBuilder {
	return qb.Update(user.Name)
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// Where appends a list comparators to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...qb.Cmp) *UserUpdate {
	uu.ub.Where(ps...)
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return uu.cqlSave(ctx)
}

func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

func (uu *UserUpdate) cqlSave(ctx context.Context) (int, error) {
	var (
		_node   = uu.createSpec()
		q       = uu.ub.QueryContext(ctx, *uu.session).BindStruct(_node)
		counter = q.Iter().NumRows()
	)
	if err := q.ExecRelease(); err != nil {
		return 0, err
	}
	return counter, nil
}

func (uu *UserUpdate) createSpec() *model.User {
	var _node = &model.User{}
	if value, ok := uu.mutation.ID(); ok {
		_node.ID = value
	}
	if value, ok := uu.mutation.Name(); ok {
		uu.ub.Set(user.FieldName)
		_node.Name = value
	}
	return _node
}

type UserUpdateOne struct {
	uu *UserUpdate
}

// SetName sets the "name" field.
func (uu *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uu.uu.mutation.SetName(s)
	return uu
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...qb.Cmp) *UserUpdateOne {
	uuo.uu.Where(ps...)
	return uuo
}

func (uuo *UserUpdateOne) Save(ctx context.Context) error {
	n, err := uuo.uu.Save(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{user.Label, *uuo.uu.mutation.id}
	default:
		return nil
	}
}
