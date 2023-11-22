package model

import (
	"context"
	"time"

	"github.com/gocql/gocql"
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

func (m *User) generateUUIFromTime() *User {
	m.ID = gocql.UUIDFromTime(time.Now()).String()
	return m
}

func (m *User) SetName(name string) *User {
	m.Name = name
	return m
}

func (m *User) GetQueryContext(ctx context.Context, columns ...string) (*User, error) {
	user := User{
		ID:   m.ID,
		Name: m.Name,
	}

	uq := m.table.GetQueryContext(ctx, *m.session, columns...).BindStruct(user)
	if err := uq.Get(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *User) InsertQueryContext(ctx context.Context) (*User, error) {
	m.generateUUIFromTime()

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
