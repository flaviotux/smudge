package model

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
	"gitlab.luizalabs.com/luizalabs/smudge/db"
)

type User struct {
	session *gocqlx.Session
	table   table.Table

	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func NewUserModel(session *gocqlx.Session) *User {
	return &User{session: session, table: *db.UserTable}
}

func (m *User) SetID(id string) *User {
	m.ID = id
	return m
}

func (m *User) SetName(name string) *User {
	m.Name = name
	return m
}

func (m *User) GetQueryContext(ctx context.Context, columns ...string) (*User, error) {
	uq := m.table.GetQueryContext(ctx, *m.session, columns...).BindStruct(m)
	if err := uq.Get(&m); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *User) InsertQueryContext(ctx context.Context) (*User, error) {
	q := m.table.InsertQueryContext(ctx, *m.session).BindStruct(m)
	if err := q.ExecRelease(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *User) SelectQueryContext(ctx context.Context, arg map[string]interface{}, columns ...string) ([]*User, error) {
	var user []*User

	q := m.table.SelectQueryContext(ctx, *m.session, columns...).BindMap(arg)
	if err := q.SelectRelease(&user); err != nil {
		return nil, err
	}

	return user, nil
}
