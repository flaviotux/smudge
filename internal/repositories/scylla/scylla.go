package scylla

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

var (
	scyllaCluster           = os.Getenv("SCYLLA_DB_CLUSTER")
	scyllaKeyspace          = os.Getenv("SCYLLA_DB_KEYSPACE")
	scyllaRF, _             = strconv.Atoi(os.Getenv("SCYLLA_DB_KEYSPACE_REPLICATION_FACTOR"))
	scyllaClusterTimeout, _ = strconv.Atoi(os.Getenv("SCYLLA_DB_CLUSTER_TIMEOUT"))
	scyllaClusterRetry, _   = strconv.Atoi(os.Getenv("SCYLLA_DB_CLUSTER_RETRY"))
)

// CreateSession creates a new gocqlx session from flags.
func CreateSession() *Session {
	cluster := CreateCluster(scyllaCluster)
	session := CreateSessionFromCluster(cluster)
	return NewSession(session)
}

// CreateCluster creates gocql ClusterConfig from flags.
func CreateCluster(clusters string) *gocql.ClusterConfig {
	clusterHosts := strings.Split(clusters, ",")

	cluster := gocql.NewCluster(clusterHosts...)
	cluster.Timeout = time.Duration(scyllaClusterTimeout) * time.Second
	cluster.Consistency = gocql.Quorum

	if scyllaClusterRetry > 0 {
		cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: scyllaClusterRetry}
	}

	cluster.Compressor = &gocql.SnappyCompressor{}

	return cluster
}

// CreateKeyspace creates keyspace with SimpleStrategy and RF derived from flags.
func CreateKeyspace(cluster *gocql.ClusterConfig, keyspace string) error {
	c := *cluster
	c.Keyspace = "system"
	c.Timeout = 30 * time.Second

	session := CreateSessionFromCluster(&c)

	{
		err := session.ExecStmt(
			fmt.Sprintf(`
				CREATE KEYSPACE %s
				WITH replication = {'class' : 'NetworkTopologyStrategy', 'replication_factor' : %d}
				AND durable_writes = true;
			`, keyspace, scyllaRF))
		if err != nil {
			return fmt.Errorf("create keyspace: %w", err)
		}
	}

	return nil
}

func CreateSessionFromCluster(cluster *gocql.ClusterConfig) gocqlx.Session {
	cluster.Keyspace = scyllaKeyspace
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal("CreateSession:", err)
	}
	defer session.Close()
	return session
}
