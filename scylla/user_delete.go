package scylla

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla/user"
)

type UserDelete struct {
	db       *qb.DeleteBuilder
	mutation *UserMutation
	session  *gocqlx.Session
}

func newUserDeleteBuilder() *qb.DeleteBuilder {
	return qb.Delete(user.Name)
}

// Where appends a list comparators to the UserDelete builder.
func (ud *UserDelete) Where(ps ...qb.Cmp) *UserDelete {
	ud.db.Where(ps...)
	return ud
}

// Exec executes the deletion query and returns how many Users were deleted.
func (ud *UserDelete) Exec(ctx context.Context) (int, error) {
	return ud.cqlExec(ctx)
}

func (ud *UserDelete) cqlExec(ctx context.Context) (int, error) {
	var (
		q       = ud.db.QueryContext(ctx, *ud.session)
		counter = q.Iter().NumRows()
	)
	if err := q.ExecRelease(); err != nil {
		return 0, err
	}
	return counter, nil
}

// UserDeleteOne is the builder for deleting a single User entity.
type UserDeleteOne struct {
	ud *UserDelete
}

// Where appends a list predicates to the UserDelete builder.
func (udo *UserDeleteOne) Where(ps ...qb.Cmp) *UserDeleteOne {
	udo.ud.Where(ps...)
	return udo
}

// Exec executes the deletion query.
func (udo *UserDeleteOne) Exec(ctx context.Context) error {
	n, err := udo.ud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{user.Label, *udo.ud.mutation.id}
	default:
		return nil
	}
}
