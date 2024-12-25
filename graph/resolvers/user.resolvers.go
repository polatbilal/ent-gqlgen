package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.61

import (
	"context"
	"fmt"
	"gqlgen-ent/ent"
	"gqlgen-ent/ent/user"
	"gqlgen-ent/graph/generated"
	"gqlgen-ent/graph/model"
	"gqlgen-ent/middlewares"
	"strconv"
	"time"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UserInput) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*ent.User, error) {
	client := middlewares.GetClientFromContext(ctx)
	intID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return client.User.Query().Where(user.IDEQ(intID)).Only(ctx)
}

// AllUsers is the resolver for the allUsers field.
func (r *queryResolver) AllUsers(ctx context.Context) ([]*ent.User, error) {
	client := middlewares.GetClientFromContext(ctx)
	return client.User.Query().All(ctx)
}

// CreatedAt is the resolver for the createdAt field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *ent.User) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
