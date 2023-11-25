package scylla

import (
	"errors"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type (
	ScyllaManager struct {
		hosts    []string
		keyspace string
	}

	ValidationError struct {
		Name string // Field or edge name.
		err  error
	}
)

func NewScyllaManager(hosts []string, keyspace string) *ScyllaManager {
	return &ScyllaManager{hosts, keyspace}
}

func (m *ScyllaManager) Connect() (gocqlx.Session, error) {
	cluster := gocql.NewCluster(m.hosts...)
	cluster.Keyspace = m.keyspace

	return gocqlx.WrapSession(cluster.CreateSession())
}

func (m *ScyllaManager) CreateKeyspace() error {
	cluster := gocql.NewCluster(m.hosts...)
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return err
	}
	defer session.Close()

	stmt := fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`, m.keyspace)
	return session.ExecStmt(stmt)
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return e.err.Error()
}

// Unwrap implements the errors.Wrapper interface.
func (e *ValidationError) Unwrap() error {
	return e.err
}

// IsValidationError returns a boolean indicating whether the error is a validation error.
func IsValidationError(err error) bool {
	if err == nil {
		return false
	}
	var e *ValidationError
	return errors.As(err, &e)
}
