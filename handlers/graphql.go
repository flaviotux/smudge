package handlers

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/scylladb/gocqlx/v2"
	"gitlab.luizalabs.com/luizalabs/smudge/graph"
)

func HandleGraphQL(session *gocqlx.Session) http.HandlerFunc {
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
