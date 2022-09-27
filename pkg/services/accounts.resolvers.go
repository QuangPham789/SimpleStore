package services

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphdemo/pkg/entity"
	"graphdemo/pkg/graph/generated"
	"graphdemo/pkg/repositories"
	"log"

	_ "github.com/lib/pq"
)

// Order is the resolver for the order field.
func (r *accountsResolver) Order(ctx context.Context, obj *entity.Accounts) ([]*entity.Order, error) {
	var m []*entity.Order

	lstOrder, err := repositories.GetAllOrder(ctx)

	if err != nil {
		log.Fatal(err)
	}
	for _, o := range lstOrder {
		if o.Accountid == obj.ID {
			m = append(m, o)
		}
	}

	return m, nil
}

// Accounts returns generated.AccountsResolver implementation.
func (r *Resolver) Accounts() generated.AccountsResolver { return &accountsResolver{r} }

type accountsResolver struct{ *Resolver }
