package json

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gitlab.luizalabs.com/luizalabs/smudge/handlers"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla"
)

type JSONAPIServer struct {
	ListenAddr string
	session    *scylla.Session
}

func MakeJSONAPIServerAndRun(listenAddr string, session *scylla.Session) error {
	server := NewJSONAPIServer(listenAddr, session)
	return server.Run()
}

func (s *JSONAPIServer) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/graphql", func(r chi.Router) {
		r.Post("/", handlers.HandleGraphQL(s.session))
		r.Get("/", handlers.HandleGraphQLPlayground())
	})

	log.Printf("Serving JSON API on 0.0.0.0%s\n", s.ListenAddr)
	return http.ListenAndServe(s.ListenAddr, r)
}

func NewJSONAPIServer(listenAddr string, session *scylla.Session) *JSONAPIServer {
	return &JSONAPIServer{listenAddr, session}
}
