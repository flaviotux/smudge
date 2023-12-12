package todo

import "github.com/scylladb/gocqlx/v2/table"

var (
	Table = table.New(table.Metadata{
		Name:    Name,
		Columns: Columns,
		PartKey: PartKey,
		SortKey: SortKey,
	})
)
