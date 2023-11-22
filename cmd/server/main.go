package main

import (
	"flag"
	"log"
	"os"

	"gitlab.luizalabs.com/luizalabs/smudge/db"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/app"
)

func main() {
	var (
		restAddr = flag.String("rest", os.Getenv("REST_LISTENADDR"), "listen address of the rest transport")
		grpcAddr = flag.String("grpc", os.Getenv("GRPC_LISTENADDR"), "listen address of the grpc transport")
		hosts    = []string{"localhost"}
		keyspace = "smudge"
	)
	flag.Parse()

	manager := db.NewScyllaManager(hosts, keyspace)

	if err := manager.CreateKeyspace(); err != nil {
		log.Fatal(err)
	}

	session, err := manager.Connect()
	if err != nil {
		log.Fatal(err)
	}

	go app.MakeGRPCServerAndRun(*grpcAddr, &session)

	restServer := app.NewRESTAPIServer(*restAddr, &session)
	restServer.Run()
}
