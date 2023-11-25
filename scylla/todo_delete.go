package scylla

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla/todo"
)

type TodoDelete struct {
	mutation *TodoMutation
	session  *gocqlx.Session
}

// Where appends a list predicates to the CarDelete builder.
func (td *TodoDelete) Where(ps ...qb.Cmp) *TodoDelete {
	td.mutation.Where(ps...)
	return td
}

// Exec executes the deletion query and returns how many todos were deleted.
func (td *TodoDelete) Exec(ctx context.Context) (int, error) {
	return td.sqlExec(ctx)
}

func (td *TodoDelete) sqlExec(ctx context.Context) (int, error) {
	var counter int
	sb := todoTable.SelectBuilder().Where(td.mutation.comparators...).CountAll()
	qs := sb.QueryContext(ctx, *td.session)
	if err := qs.SelectRelease(&counter); err != nil {
		return 0, err
	}
	db := todoTable.DeleteBuilder().Existing().Where(td.mutation.comparators...)
	qd := db.QueryContext(ctx, *td.session)
	if err := qd.ExecRelease(); err != nil {
		return 0, err
	}
	return counter, nil
}

// TodoDeleteOne is the builder for deleting a single Todo entity.
type TodoDeleteOne struct {
	td *TodoDelete
}

// Where appends a list predicates to the CarDelete builder.
func (tdo *TodoDeleteOne) Where(ps ...qb.Cmp) *TodoDeleteOne {
	tdo.td.mutation.Where(ps...)
	return tdo
}

// Exec executes the deletion query.
func (tdo *TodoDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{todo.Label, *tdo.td.mutation.id}
	default:
		return nil
	}
}
