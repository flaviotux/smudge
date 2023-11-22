package db

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type ScyllaManager struct {
	hosts    []string
	keyspace string
}

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
