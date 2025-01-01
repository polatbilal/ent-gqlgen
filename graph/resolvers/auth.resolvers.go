package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.61

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/polatbilal/gqlgen-ent/database"
	"github.com/polatbilal/gqlgen-ent/ent"
	"github.com/polatbilal/gqlgen-ent/ent/user"
	"github.com/polatbilal/gqlgen-ent/graph/generated"
	"github.com/polatbilal/gqlgen-ent/graph/model"
	"github.com/polatbilal/gqlgen-ent/middlewares"
	"github.com/polatbilal/gqlgen-ent/services"
	"github.com/polatbilal/gqlgen-ent/tools"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, companyCode string, username string, name *string, email *string, password string) (*model.AuthPayload, error) {
	client := middlewares.GetClientFromContext(ctx)
	_, err := client.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err == nil {
		return nil, fmt.Errorf("kullanıcı adı zaten mevcut")
	}

	if !ent.IsNotFound(err) {
		return nil, err
	}

	user, err := client.User.Create().
		SetUsername(username).
		SetPassword(tools.HashPassword(password)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	var userID = strconv.Itoa(user.ID)

	token, err := services.JwtGenerate(ctx, user.ID, username, user.Name, companyCode, user.Role)
	if err != nil {
		return nil, err
	}

	return &model.AuthPayload{
		Token:       token,
		UserID:      userID,
		Username:    user.Username,
		Name:        user.Name,
		CompanyCode: companyCode,
		Role:        user.Role,
	}, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, companyCode string, username string, password string) (*model.AuthPayload, error) {
	client, err := database.GetClient(companyCode)
	if err != nil {
		log.Printf("Failed to get DB client: %v", err)
		return nil, err
	}
	user, err := client.User.Query().
		Where(user.UsernameEQ(username)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	if err := tools.ComparePassword(user.Password, password); err != nil {
		return nil, err
	}

	var userID = strconv.Itoa(user.ID)

	token, err := services.JwtGenerate(ctx, user.ID, username, user.Name, companyCode, user.Role)
	if err != nil {
		return nil, err
	}

	// Token'ı Redis'e kaydet (24 saat geçerli)
	err = database.RedisClient.Set(ctx, token, userID, 24*time.Hour).Err()
	if err != nil {
		return nil, fmt.Errorf("token kaydedilemedi: %v", err)
	}

	return &model.AuthPayload{
		Token:       token,
		UserID:      userID,
		Username:    username,
		Name:        user.Name,
		CompanyCode: companyCode,
		Role:        user.Role,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
