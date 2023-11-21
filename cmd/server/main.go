package main

import (
	"flag"
	"os"

	"gitlab.luizalabs.com/luizalabs/smudge/internal/app"
)

func main() {
	var (
		restAddr = flag.String("rest", os.Getenv("REST_LISTENADDR"), "listen address of the rest transport")
		grpcAddr = flag.String("grpc", os.Getenv("GRPC_LISTENADDR"), "listen address of the grpc transport")
	)
	flag.Parse()

	go app.MakeGRPCServerAndRun(*grpcAddr)

	restServer := app.NewRESTAPIServer(*restAddr)
	restServer.Run()
}
