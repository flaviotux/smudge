package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gitlab.luizalabs.com/luizalabs/smudge/graph"
	"gitlab.luizalabs.com/luizalabs/smudge/pb"
	"gitlab.luizalabs.com/luizalabs/smudge/smudge"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultPort = "8080"

func graphqlHandler() http.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func playgroundHandler() http.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalln("failed to create listener:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterTodoServer(s, &smudge.Server{})

	go func() {
		log.Println("Serving gRPC on 0.0.0.0:8090")
		if err := s.Serve(listener); err != nil {
			log.Fatalln("failed to serve:", err)
		}
	}()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/graphql", graphqlHandler())
	r.Get("/graphql", playgroundHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
