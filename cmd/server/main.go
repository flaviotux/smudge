package main

import (
	"os"

	"gitlab.luizalabs.com/luizalabs/smudge/internal/repositories/scylla"
)

var (
	listenaddrGRPC = os.Getenv("REST_LISTENADDR")
	listenaddrRest = os.Getenv("GRPC_LISTENADDR")
)

func main() {
	session := scylla.CreateSession()

	go makeJSONAPIServerAndRun(listenaddrRest, session)
	makeGRPCServerAndRun(listenaddrGRPC, session)
}
