package repositories

import "gitlab.luizalabs.com/luizalabs/smudge/internal/repositories/scylla"

func NewScyllaDB() *scylla.Session {
	return scylla.CreateSession()
}
