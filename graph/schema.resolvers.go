package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"github.com/danyukod/chave-pix-bff/graph/model"
	"github.com/danyukod/chave-pix-bff/internal/application/commands/dto"
)

// CreatePixKey is the resolver for the createPixKey field.
func (r *mutationResolver) CreatePixKey(ctx context.Context, newPixKey model.NewPixKey) (*model.PixKey, error) {
	createPixKeyInputDTO := &dto.CreatePixKeyInputDTO{
		PixKeyType:            newPixKey.PixKeyType,
		PixKey:                newPixKey.PixKey,
		AccountType:           newPixKey.AccountType,
		AgencyNumber:          newPixKey.AgencyNumber,
		AccountNumber:         newPixKey.AccountNumber,
		AccountHolderName:     newPixKey.HolderName,
		AccountHolderLastName: *newPixKey.HolderLastName,
	}

	domain, err := r.CreatePixKeyUsecase.Execute(createPixKeyInputDTO)
	if err != nil {
		return nil, err
	}

	return &model.PixKey{
		ID:             domain.GetID(),
		PixKeyType:     domain.GetPixKeyType().GetType(),
		PixKey:         domain.GetPixKey(),
		AccountType:    domain.GetAccount().GetAccountType().String(),
		AgencyNumber:   domain.GetAccount().GetAgency(),
		AccountNumber:  domain.GetAccount().GetNumber(),
		HolderName:     domain.GetAccount().GetHolder().GetName(),
		HolderLastName: domain.GetAccount().GetHolder().GetLastName(),
	}, nil
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
	pixKeyDto, err := r.FindPixKeyUsecase.Execute()
	if err != nil {
		return nil, err
	}
	var response []*model.PixKey
	for _, pixKey := range *pixKeyDto {
		response = append(response, &model.PixKey{
			ID:             pixKey.Id,
			PixKeyType:     pixKey.PixKeyType,
			PixKey:         pixKey.PixKey,
			AccountType:    pixKey.AccountType,
			AgencyNumber:   pixKey.AgencyNumber,
			AccountNumber:  pixKey.AccountNumber,
			HolderName:     pixKey.AccountHolderName,
			HolderLastName: pixKey.AccountHolderLastName,
		})
	}

	return response, nil
}

// PixKey is the resolver for the pixKey field.
func (r *queryResolver) PixKey(ctx context.Context, id string) (*model.PixKey, error) {
	findPixKeyByKeyDTO, err := r.FindPixKeyByKeyUsecase.Execute(id)
	if err != nil {
		return nil, err
	}
	return &model.PixKey{
		ID:             findPixKeyByKeyDTO.Id,
		PixKeyType:     findPixKeyByKeyDTO.PixKeyType,
		PixKey:         findPixKeyByKeyDTO.PixKey,
		AccountType:    findPixKeyByKeyDTO.AccountType,
		AgencyNumber:   findPixKeyByKeyDTO.AgencyNumber,
		AccountNumber:  findPixKeyByKeyDTO.AccountNumber,
		HolderName:     findPixKeyByKeyDTO.AccountHolderName,
		HolderLastName: findPixKeyByKeyDTO.AccountHolderLastName,
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
