package main

import (
	"os"

	"gitlab.luizalabs.com/luizalabs/smudge/internal/grpc"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/rest"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla"
)

var (
	listenaddrGRPC = os.Getenv("REST_LISTENADDR")
	listenaddrRest = os.Getenv("GRPC_LISTENADDR")
)

func main() {
	session := scylla.CreateSession()

	go rest.MakeRESTAPIServerAndRun(listenaddrRest, session)
	grpc.MakeGRPCServerAndRun(listenaddrGRPC, session)
}
