package resolvers

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/polatbilal/ent-gqlgen/directives"
	"github.com/polatbilal/ent-gqlgen/ent"
	gqlutils "github.com/polatbilal/ent-gqlgen/graphql"
	"github.com/polatbilal/ent-gqlgen/graphql/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client         *ent.Client
	batchSemaphore *gqlutils.Semaphore
}

func NewResolver(client *ent.Client) *Resolver {
	return &Resolver{
		client:         client,
		batchSemaphore: gqlutils.NewSemaphore(10), // Max 10 eşzamanlı batch işlem
	}
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
