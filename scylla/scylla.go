package scylla

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

var (
	scyllaCluster  = os.Getenv("SCYLLA_DB_CLUSTER")
	scyllaKeyspace = os.Getenv("SCYLLA_DB_KEYSPACE")

	flagProto   = flag.Int("proto", 0, "protcol version")
	flagCQL     = flag.String("cql", "3.0.0", "CQL version")
	flagRF      = flag.Int("rf", 1, "replication factor for test keyspace")
	flagRetry   = flag.Int("retries", 5, "number of times to retry queries")
	flagTimeout = flag.Duration("gocql.timeout", 5*time.Second, "sets the connection `timeout` for all operations")
)

type (
	ValidationError struct {
		Name string // Field or edge name.
		err  error
	}
	// NotFoundError returns when trying to update an
	// entity, and it was not found in the database.
	NotFoundError struct {
		table string
		id    string
	}
)

// CreateSession creates a new gocqlx session from flags.
func CreateSession() *Session {
	cluster := CreateCluster()
	session := createSessionFromCluster(cluster)
	return NewSession(session)
}

// CreateCluster creates gocql ClusterConfig from flags.
func CreateCluster() *gocql.ClusterConfig {
	if !flag.Parsed() {
		flag.Parse()
	}
	clusterHosts := strings.Split(scyllaCluster, ",")

	cluster := gocql.NewCluster(clusterHosts...)
	cluster.ProtoVersion = *flagProto
	cluster.CQLVersion = *flagCQL
	cluster.Timeout = *flagTimeout
	cluster.Consistency = gocql.Quorum
	cluster.MaxWaitSchemaAgreement = 2 * time.Minute // travis might be slow
	if *flagRetry > 0 {
		cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: *flagRetry}
	}

	cluster.Compressor = &gocql.SnappyCompressor{}

	return cluster
}

// CreateKeyspace creates keyspace with SimpleStrategy and RF derived from flags.
func CreateKeyspace(cluster *gocql.ClusterConfig, keyspace string) error {
	c := *cluster
	c.Keyspace = "system"
	c.Timeout = 30 * time.Second

	session, err := gocqlx.WrapSession(c.CreateSession())
	if err != nil {
		return err
	}
	defer session.Close()

	{
		err := session.ExecStmt(
			fmt.Sprintf(`
				CREATE KEYSPACE %s
				WITH replication = {'class' : 'NetworkTopologyStrategy', 'replication_factor' : %d}
				AND durable_writes = true;
			`, keyspace, *flagRF))
		if err != nil {
			return fmt.Errorf("create keyspace: %w", err)
		}
	}

	return nil
}

func createSessionFromCluster(cluster *gocql.ClusterConfig) gocqlx.Session {
	if !flag.Parsed() {
		flag.Parse()
	}

	cluster.Keyspace = scyllaKeyspace
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal("CreateSession:", err)
	}
	return session
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

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("record with id %v not found in table %s", e.id, e.table)
}
