package database

import (
	"fmt"
	"github.com/gocql/gocql"
)

const (
	CASSANDRA_USERNAME     = "cassandra"
	CASSANDRA_PASSWORD     = "cassandra"
	CASSANDRA_CLUSTER_NAME = "noty"
	CASSANDRA_HOST         = "127.0.0.1"
	CASSANDRA_PORT         = 9042
	CASSANDRA_KEYSPACE     = "noty"
)

func InitCassandra() *gocql.Session {
	cluster := gocql.NewCluster(CASSANDRA_HOST)
	cluster.Keyspace = CASSANDRA_KEYSPACE
	cluster.Consistency = gocql.Any
	cluster.Port = CASSANDRA_PORT
	cluster.Hosts = []string{CASSANDRA_HOST}

	session, err := cluster.CreateSession()

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return session
}
