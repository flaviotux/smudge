package db

import "github.com/scylladb/gocqlx/v2/table"

var userMetadata = table.Metadata{
	Name: "users",
	Columns: []string{
		"id", "name",
	},
	PartKey: []string{},
	SortKey: []string{"id"},
}

var UserTable = table.New(userMetadata)
