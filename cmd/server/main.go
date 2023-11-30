package main

import (
	"os"

	"gitlab.luizalabs.com/luizalabs/smudge/internal/app"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla"
)

var (
	listenaddrGRPC = os.Getenv("REST_LISTENADDR")
	listenaddrRest = os.Getenv("GRPC_LISTENADDR")
)

func main() {
	session := scylla.CreateSession()

	go app.MakeRESTAPIServerAndRun(listenaddrRest, session)
	app.MakeGRPCServerAndRun(listenaddrGRPC, session)
}
