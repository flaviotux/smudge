package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/handlers"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/repositories/scylla"
)

type JSONAPIServer struct {
	listenAddr string
	session    *scylla.Session
}

func makeJSONAPIServerAndRun(listenAddr string, session *scylla.Session) error {
	server := newJSONAPIServer(listenAddr, session)
	return server.run()
}

func (s *JSONAPIServer) run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/graphql", func(r chi.Router) {
		r.Post("/", handlers.HandleGraphQL(s.session))
		r.Get("/", handlers.HandleGraphQLPlayground())
	})

	log.Printf("Serving JSON API on 0.0.0.0%s\n", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, r)
}

func newJSONAPIServer(listenAddr string, session *scylla.Session) *JSONAPIServer {
	return &JSONAPIServer{listenAddr, session}
}
