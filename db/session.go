package db

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/model"
)

type Session struct {
	gocqlx.Session
	Todo *TodoSession
	User *UserSession
}

// NewSession creates a new client configured with the given options.
func NewSession(session gocqlx.Session) *Session {
	s := &Session{Session: session}
	s.init()
	return s
}

func (s *Session) init() {
	s.Todo = NewTodoSession(&s.Session)
	s.User = NewUserSession(&s.Session)
}

type TodoSession struct {
	session *gocqlx.Session
}

// NewTodoSession returns a client for the Todo from the given session.
func NewTodoSession(session *gocqlx.Session) *TodoSession {
	return &TodoSession{session}
}

// Create returns a builder for creating a Todo entity.
func (s *TodoSession) Create() *TodoCreate {
	mutation := newTodoMutation()
	return &TodoCreate{
		mutation: mutation,
		session:  s.session,
		Table:    *TodoTable,
	}
}

// Query returns a query builder for Todo.
func (s *TodoSession) Query() *TodoQuery {
	return &TodoQuery{
		session: s.session,
		Table:   *TodoTable,
	}
}

type UserSession struct {
	session *gocqlx.Session
}

// NewUserSession returns a client for the User from the given session.
func NewUserSession(session *gocqlx.Session) *UserSession {
	return &UserSession{session}
}

// Create returns a builder for creating a User entity.
func (s *UserSession) Create() *UserCreate {
	mutation := newUserMutation()
	return &UserCreate{
		mutation: mutation,
		session:  s.session,
		Table:    *UserTable,
	}
}

// Query returns a query builder for User.
func (s *UserSession) Query() *UserQuery {
	return &UserQuery{
		session: s.session,
		Table:   *UserTable,
	}
}

// Get returns a Car entity by its id.
func (s *UserSession) Get(ctx context.Context, id string) (*model.User, error) {
	return s.Query().Where(qb.EqLit("id", id)).Only(ctx)
}
