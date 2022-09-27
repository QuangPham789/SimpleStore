package services

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphdemo/pkg/entity"
	"graphdemo/pkg/graph/generated"
	"graphdemo/pkg/repositories"
)

// Order is the resolver for the order field.
func (r *orderResolver) Order(ctx context.Context, obj *entity.Order) (*entity.Order, error) {
	return repositories.For(ctx).OrderLoader.Load(obj.ID)
}

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

type orderResolver struct{ *Resolver }
