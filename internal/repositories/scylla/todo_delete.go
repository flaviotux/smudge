package scylla

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/repositories/scylla/todo"
)

type TodoDelete struct {
	db       *qb.DeleteBuilder
	mutation *TodoMutation
	session  *gocqlx.Session
}

func newTodoDeleteBuilder() *qb.DeleteBuilder {
	return qb.Delete(todo.Name)
}

// Where appends a list comparators to the TodoDelete builder.
func (td *TodoDelete) Where(ps ...qb.Cmp) *TodoDelete {
	td.db.Where(ps...)
	return td
}

// Exec executes the deletion query and returns how many todos were deleted.
func (td *TodoDelete) Exec(ctx context.Context) (int, error) {
	return td.cqlExec(ctx)
}

func (td *TodoDelete) cqlExec(ctx context.Context) (int, error) {
	var (
		q       = td.db.QueryContext(ctx, *td.session)
		counter = q.Iter().NumRows()
	)
	if err := q.ExecRelease(); err != nil {
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
	tdo.td.Where(ps...)
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
