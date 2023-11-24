package db

type TodoMutation struct {
	id      *string
	text    *string
	done    *bool
	user_id *string
}

// newTodoMutation creates new mutation for the Todo entity.
func newTodoMutation() *TodoMutation {
	return &TodoMutation{}
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TodoMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetText sets the "text" field.
func (m *TodoMutation) SetText(s string) {
	m.text = &s
}

// Text returns the value of the "text" field in the mutation.
func (m *TodoMutation) Text() (r string, exists bool) {
	v := m.text
	if v == nil {
		return
	}
	return *v, true
}

// SetText sets the "text" field.
func (m *TodoMutation) SetDone(b bool) {
	m.done = &b
}

// Done returns the value of the "done" field in the mutation.
func (m *TodoMutation) Done() (r bool, exists bool) {
	v := m.done
	if v == nil {
		return
	}
	return *v, true
}

// SetText sets the "text" field.
func (m *TodoMutation) SetUserId(s string) {
	m.user_id = &s
}

// Done returns the value of the "done" field in the mutation.
func (m *TodoMutation) UserID() (r string, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

type UserMutation struct {
	id   *string
	name *string
}

// newUserMutation creates new mutation for the Todo entity.
func newUserMutation() *UserMutation {
	return &UserMutation{}
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "Name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "Name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}
