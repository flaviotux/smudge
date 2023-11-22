package model

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
	"gitlab.luizalabs.com/luizalabs/smudge/db"
)

type Todo struct {
	session *gocqlx.Session
	table   table.Table

	ID     string `json:"id" db:"id"`
	Text   string `json:"text" db:"text"`
	Done   bool   `json:"done" db:"done"`
	UserID string `json:"userId" db:"user_id"`
	User   *User
}

func NewTodoModel(session *gocqlx.Session) *Todo {
	return &Todo{session: session, table: *db.TodoTable}
}

func (m *Todo) SetID(id string) *Todo {
	m.ID = id
	return m
}

func (m *Todo) SetText(text string) *Todo {
	m.Text = text
	return m
}

func (m *Todo) SetDone(done bool) *Todo {
	m.Done = done
	return m
}

func (m *Todo) SetUserID(userID string) *Todo {
	m.UserID = userID
	return m
}

func (m *Todo) SetUser(user *User) *Todo {
	m.User = user
	return m
}

func (m *Todo) GetQueryContext(ctx context.Context, columns ...string) (*Todo, error) {
	uq := m.table.GetQueryContext(ctx, *m.session, columns...).BindStruct(m)
	if err := uq.Get(&m); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Todo) InsertQueryContext(ctx context.Context) (*Todo, error) {
	q := m.table.InsertQueryContext(ctx, *m.session).BindStruct(m)
	if err := q.ExecRelease(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Todo) SelectQueryContext(ctx context.Context, arg map[string]interface{}, columns ...string) ([]*Todo, error) {
	var todo []*Todo

	q := m.table.SelectQueryContext(ctx, *m.session).BindMap(arg)
	if err := q.SelectRelease(&todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (m *Todo) WithUserContext(ctx context.Context, columns ...string) (*Todo, error) {
	u := NewUserModel(m.session).SetID(m.UserID)

	if _, err := u.GetQueryContext(ctx, columns...); err != nil {
		return nil, err
	}

	m.SetUser(u)

	return m, nil
}
