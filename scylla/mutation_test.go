package scylla_test

import (
	"testing"

	"github.com/gocql/gocql"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla"
)

func TestTodoMutation(t *testing.T) {
	mutation := scylla.TodoMutation{}
	t.Run("!ID", func(t *testing.T) {
		if _, ok := mutation.ID(); ok {
			t.Fail()
		}
	})
	t.Run("!Text", func(t *testing.T) {
		if _, ok := mutation.Text(); ok {
			t.Fail()
		}
	})
	t.Run("Text", func(t *testing.T) {
		mutation.SetText("todo")
		if _, ok := mutation.Text(); !ok {
			t.Fail()
		}
	})
	t.Run("!UserID", func(t *testing.T) {
		if _, ok := mutation.UserID(); ok {
			t.Fail()
		}
	})
	t.Run("UserID", func(t *testing.T) {
		mutation.SetUserId(gocql.TimeUUID().String())
		if _, ok := mutation.UserID(); !ok {
			t.Fail()
		}
	})
}

func TestUserMutation(t *testing.T) {
	mutation := scylla.UserMutation{}
	t.Run("!ID", func(t *testing.T) {
		if _, ok := mutation.ID(); ok {
			t.Fail()
		}
	})
	t.Run("!Name", func(t *testing.T) {
		if _, ok := mutation.Name(); ok {
			t.Fail()
		}
	})
	t.Run("Name", func(t *testing.T) {
		mutation.SetName("user")
		if _, ok := mutation.Name(); !ok {
			t.Fail()
		}
	})
}
