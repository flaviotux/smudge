package scylla

import "github.com/scylladb/gocqlx/v2/table"

var todoMetadata = table.Metadata{
	Name: "todos",
	Columns: []string{
		"id", "text", "done", "user_id",
	},
	PartKey: []string{},
	SortKey: []string{"id"},
}

var TodoTable = table.New(todoMetadata)
