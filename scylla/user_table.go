package scylla

import (
	"github.com/scylladb/gocqlx/v2/table"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla/user"
)

var userMetadata = table.Metadata{
	Name:    user.Table,
	Columns: user.Columns,
	PartKey: user.PartKey,
	SortKey: user.SortKey,
}

var userTable = table.New(userMetadata)
