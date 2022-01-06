package graph

///o:generate go run github.com/blennster/gql-go-todo gen

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/blennster/gql-go-todo/graph/generated"

	"github.com/99designs/gqlgen/graphql"
	"github.com/samonzeweb/godb"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db *godb.DB
}

func NewSchema(db *godb.DB) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{db: db},
	})
}
