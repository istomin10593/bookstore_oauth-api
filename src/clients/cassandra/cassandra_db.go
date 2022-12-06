package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/istomin10593/bookstore_oauth-api/src/logger"
)

var (
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error

	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}

	logger.Info("cassandra connection successfully created")

}

func GetSession() *gocql.Session {
	return session
}
