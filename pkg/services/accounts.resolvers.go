package services

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphdemo/pkg/entity"
	"graphdemo/pkg/graph/generated"
	"graphdemo/pkg/graph/model"
	"graphdemo/pkg/repositories"

	_ "github.com/lib/pq"
)

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, input model.NewAccount) (*entity.Accounts, error) {
	accountResult, err := repositories.CreateAccount(ctx, input)
	if err != nil {
		return nil, err
	}
	return accountResult, nil
}

// UpdateAccount is the resolver for the updateAccount field.
func (r *mutationResolver) UpdateAccount(ctx context.Context, input model.UpdateAccountModel) (*entity.Accounts, error) {
	accountResult, err := repositories.UpdateAccount(ctx, input)
	if err != nil {
		return nil, err
	}
	return accountResult, nil
}

// DeleteAccount is the resolver for the deleteAccount field.
func (r *mutationResolver) DeleteAccount(ctx context.Context, id int) (*entity.Accounts, error) {
	panic(fmt.Errorf("not implemented: DeleteAccount - deleteAccount"))
}

// GetAllAccounts is the resolver for the GetAllAccounts field.
func (r *queryResolver) GetAllAccounts(ctx context.Context) ([]*entity.Accounts, error) {
	accountsFromDB, err := repositories.GetAllAccounts(ctx)
	if err != nil {
		return nil, err
	}

	return accountsFromDB, nil
}

// GetAccountByID is the resolver for the GetAccountByID field.
func (r *queryResolver) GetAccountByID(ctx context.Context, id int) (*entity.Accounts, error) {
	accountsFromDB, err := repositories.GetAccountByID(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	return accountsFromDB, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
