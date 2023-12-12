package scylla

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"gitlab.luizalabs.com/luizalabs/smudge/graph/model"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla/todo"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla/user"
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
	builder := newTodoInsertBuilder()
	return &TodoCreate{builder, mutation, s.session}
}

// Update returns a builder for updating a User entity.
func (s *TodoSession) Update() *TodoUpdate {
	mutation := newTodoMutation()
	builder := newTodoUpdateBuilder()
	return &TodoUpdate{builder, mutation, s.session}
}

// UpdateOne returns a builder for deleting the given entity.
func (s *TodoSession) UpdateOne(user *model.User) *TodoUpdateOne {
	return s.UpdateOneID(user.ID)
}

// UpdateOneID returns a builder for deleting the given entity by its id.
func (s *TodoSession) UpdateOneID(id string) *TodoUpdateOne {
	builder := s.Update().Where(user.ID(id))
	builder.mutation.id = &id
	return &TodoUpdateOne{builder}
}

// Query returns a query builder for Todo.
func (s *TodoSession) Query(todo *model.Todo) *TodoQuery {
	builder := newTodoSelectBuilder()
	mutation := &TodoMutation{
		id:      &todo.ID,
		text:    &todo.Text,
		done:    &todo.Done,
		user_id: &todo.UserID,
	}
	return &TodoQuery{builder, mutation, s.session}
}

// Get returns a Todo entity by its id.
func (s *TodoSession) Get(ctx context.Context, id string) (*model.Todo, error) {
	return s.Query(&model.Todo{ID: id}).Only(ctx)
}

// Delete returns a delete builder for Todo.
func (s *TodoSession) Delete() *TodoDelete {
	mutation := newTodoMutation()
	builder := newTodoDeleteBuilder()
	return &TodoDelete{builder, mutation, s.session}
}

// DeleteOne returns a builder for deleting the given entity.
func (s *TodoSession) DeleteOne(todo *model.Todo) *TodoDeleteOne {
	return s.DeleteOneID(todo.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (s *TodoSession) DeleteOneID(id string) *TodoDeleteOne {
	builder := s.Delete().Where(todo.ID(id))
	builder.mutation.id = &id
	return &TodoDeleteOne{builder}
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
	builder := newUserInsertBuilder()
	return &UserCreate{builder, mutation, s.session}
}

// Update returns a builder for updating a User entity.
func (s *UserSession) Update() *UserUpdate {
	mutation := newUserMutation()
	builder := newUserUpdateBuilder()
	return &UserUpdate{builder, mutation, s.session}
}

// UpdateOne returns a builder for deleting the given entity.
func (s *UserSession) UpdateOne(user *model.User) *UserUpdateOne {
	return s.UpdateOneID(user.ID)
}

// UpdateOneID returns a builder for deleting the given entity by its id.
func (s *UserSession) UpdateOneID(id string) *UserUpdateOne {
	builder := s.Update().Where(user.ID(id))
	builder.mutation.id = &id
	return &UserUpdateOne{builder}
}

// Query returns a query builder for User.
func (s *UserSession) Query() *UserQuery {
	builder := newUserSelectBuilder()
	return &UserQuery{builder, s.session}
}

// Get returns a User entity by its id.
func (s *UserSession) Get(ctx context.Context, id string) (*model.User, error) {
	return s.Query().Where(user.ID(id)).Only(ctx)
}

// Delete returns a delete builder for User.
func (s *UserSession) Delete() *UserDelete {
	mutation := newUserMutation()
	builder := newUserDeleteBuilder()
	return &UserDelete{builder, mutation, s.session}
}

// DeleteOne returns a builder for deleting the given entity.
func (s *UserSession) DeleteOne(user *model.User) *UserDeleteOne {
	return s.DeleteOneID(user.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (s *UserSession) DeleteOneID(id string) *UserDeleteOne {
	builder := s.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	return &UserDeleteOne{builder}
}
