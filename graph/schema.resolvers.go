package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphdemo/db/models"
	"graphdemo/graph/generated"
	"graphdemo/graph/model"
	"graphdemo/util"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, input model.NewAccount) (*model.Accounts, error) {
	panic(fmt.Errorf("not implemented: CreateAccount - createAccount"))
}

// GetAllAccounts is the resolver for the GetAllAccounts field.
func (r *queryResolver) GetAllAccounts(ctx context.Context) ([]*model.Accounts, error) {
	var accounts []*model.Accounts
	db := util.InitDB()
	accountsFromDB, err := models.Accounts().All(ctx, db)
	if err != nil {
		return nil, err
	}
	for _, account := range accountsFromDB {
		accountModel := &model.Accounts{
			ID:       int(account.ID),
			Username: account.Username,
			Password: int(account.Password),
		}
		accounts = append(accounts, accountModel)
	}

	return accounts, nil
}

// GetAccountByID is the resolver for the GetAccountByID field.
func (r *queryResolver) GetAccountByID(ctx context.Context, id int) (*model.Accounts, error) {
	var account *model.Accounts
	db := util.InitDB()
	accountsFromDB, err := models.Accounts(qm.Where("id=?", id)).One(ctx, db)
	if err != nil {
		return nil, err
	}

	account = &model.Accounts{
		ID:       int(accountsFromDB.ID),
		Username: accountsFromDB.Username,
		Password: int(accountsFromDB.Password),
	}

	return account, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
