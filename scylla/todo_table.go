package scylla

import (
	"github.com/scylladb/gocqlx/v2/table"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla/todo"
)

var todoMetadata = table.Metadata{
	Name:    todo.Table,
	Columns: todo.Columns,
	PartKey: todo.PartKey,
	SortKey: todo.SortKey,
}

var todoTable = table.New(todoMetadata)
