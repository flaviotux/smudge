package rest

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gitlab.luizalabs.com/luizalabs/smudge/handlers"
	"gitlab.luizalabs.com/luizalabs/smudge/scylla"
)

type RESTAPIServer struct {
	ListenAddr string
	session    *scylla.Session
}

func MakeRESTAPIServerAndRun(listenAddr string, session *scylla.Session) error {
	server := NewRESTAPIServer(listenAddr, session)
	return server.Run()
}

func (s *RESTAPIServer) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/graphql", func(r chi.Router) {
		r.Post("/", handlers.HandleGraphQL(s.session))
		r.Get("/", handlers.HandleGraphQLPlayground())
	})

	log.Printf("Serving REST API on 0.0.0.0%s\n", s.ListenAddr)
	return http.ListenAndServe(s.ListenAddr, r)
}

func NewRESTAPIServer(listenAddr string, session *scylla.Session) *RESTAPIServer {
	return &RESTAPIServer{listenAddr, session}
}
