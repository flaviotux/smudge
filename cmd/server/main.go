package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"gitlab.luizalabs.com/luizalabs/smudge/internal/app"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla"
)

func main() {
	var (
		restAddr = flag.String("rest", os.Getenv("REST_LISTENADDR"), "listen address of the rest transport")
		grpcAddr = flag.String("grpc", os.Getenv("GRPC_LISTENADDR"), "listen address of the grpc transport")
		keyspace = flag.String("keyspaces", os.Getenv("SCYLLA_DB_KEYSPACE"), "scylladb keyspace")
		hosts    = flag.String("hosts", os.Getenv("SCYLLA_DB_HOSTS"), "scylladb hosts")
	)
	flag.Parse()

	manager := scylla.NewScyllaManager(strings.Split(*hosts, ","), *keyspace)
	if err := manager.CreateKeyspace(); err != nil {
		log.Fatal(err)
	}

	scyllaSession, err := manager.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer scyllaSession.Close()

	session := scylla.NewSession(scyllaSession)

	go app.MakeGRPCServerAndRun(*grpcAddr, session)

	restServer := app.NewRESTAPIServer(*restAddr, session)
	restServer.Run()
}
