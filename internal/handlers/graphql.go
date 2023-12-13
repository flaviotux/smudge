package handlers

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gitlab.luizalabs.com/luizalabs/smudge/graph"
	"gitlab.luizalabs.com/luizalabs/smudge/internal/repositories/scylla"
)

func HandleGraphQL(session *scylla.Session) http.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewSchema(session))

	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func HandleGraphQLPlayground() http.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}
