package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/danyukod/chave-pix-bff/graph/model"
)

// CreatePixKey is the resolver for the createPixKey field.
func (r *mutationResolver) CreatePixKey(ctx context.Context, newPixKey model.NewPixKey) (*model.PixKey, error) {
	panic(fmt.Errorf("not implemented: CreatePixKey - createPixKey"))
}

// UpdatePixKey is the resolver for the updatePixKey field.
func (r *mutationResolver) UpdatePixKey(ctx context.Context, id string, newPixKey model.NewPixKey) (*model.PixKey, error) {
	panic(fmt.Errorf("not implemented: UpdatePixKey - updatePixKey"))
}

// DeletePixKey is the resolver for the deletePixKey field.
func (r *mutationResolver) DeletePixKey(ctx context.Context, id string) (*model.PixKey, error) {
	panic(fmt.Errorf("not implemented: DeletePixKey - deletePixKey"))
}

// PixKeys is the resolver for the pixKeys field.
func (r *queryResolver) PixKeys(ctx context.Context) ([]*model.PixKey, error) {
	panic(fmt.Errorf("not implemented: PixKeys - pixKeys"))
}

// PixKey is the resolver for the pixKey field.
func (r *queryResolver) PixKey(ctx context.Context, id string) (*model.PixKey, error) {
	panic(fmt.Errorf("not implemented: PixKey - pixKey"))
}

// Account is the resolver for the account field.
func (r *queryResolver) Account(ctx context.Context, id string) (*model.Account, error) {
	panic(fmt.Errorf("not implemented: Account - account"))
}

// Holder is the resolver for the holder field.
func (r *queryResolver) Holder(ctx context.Context, id string) (*model.Holder, error) {
	panic(fmt.Errorf("not implemented: Holder - holder"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
