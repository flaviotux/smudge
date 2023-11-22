package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/scylladb/gocqlx/v2"
	"gitlab.luizalabs.com/luizalabs/smudge/handlers"
)

const defaultRestAddr = ":8080"

type RESTAPIServer struct {
	ListenAddr string
	session    *gocqlx.Session
}

func (s *RESTAPIServer) Run() error {
	if s.ListenAddr == "" {
		s.ListenAddr = defaultRestAddr
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/graphql", func(r chi.Router) {
		r.Post("/", handlers.HandleGraphQL(s.session))
		r.Get("/", handlers.HandleGraphQLPlayground())
	})

	log.Printf("Serving REST API on 0.0.0.0%s\n", s.ListenAddr)
	return http.ListenAndServe(s.ListenAddr, r)
}

func NewRESTAPIServer(listenAddr string, session *gocqlx.Session) *RESTAPIServer {
	return &RESTAPIServer{listenAddr, session}
}
