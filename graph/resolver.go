package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/99designs/gqlgen/graphql"
	"gitlab.luizalabs.com/luizalabs/smudge/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	session *db.Session
}

func NewSchema(session *db.Session) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{session},
	})
}
