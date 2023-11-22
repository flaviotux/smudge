package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/scylladb/gocqlx/v2"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	session *gocqlx.Session
}

func NewSchema(session *gocqlx.Session) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{session},
	})
}
