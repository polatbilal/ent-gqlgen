package resolvers

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/polatbilal/ent-gqlgen/directives"
	"github.com/polatbilal/ent-gqlgen/ent"
	"github.com/polatbilal/ent-gqlgen/graphql/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ client *ent.Client }

func NewResolver(client *ent.Client) *Resolver {
	return &Resolver{client: client}
}

// NewSchema işlevini güncelleyerek, çözücünüzü yaratıcıdan çözücünüze yöneltin.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	resolver := NewResolver(client)
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
		Directives: generated.DirectiveRoot{
			Auth: directives.Auth,
		},
	})
}
