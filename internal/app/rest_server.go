package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gitlab.luizalabs.com/luizalabs/smudge/handlers"
)

const defaultRestAddr = ":8080"

type RESTAPIServer struct {
	listenAddr string
}

func (s *RESTAPIServer) Run() error {
	if s.listenAddr == "" {
		s.listenAddr = defaultRestAddr
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/graphql", func(r chi.Router) {
		r.Post("/", handlers.HandleGraphQL())
		r.Get("/", handlers.HandleGraphQLPlayground())
	})

	log.Printf("Serving REST API on 0.0.0.0%s\n", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, r)
}

func NewRESTAPIServer(listenAddr string) *RESTAPIServer {
	return &RESTAPIServer{listenAddr}
}
