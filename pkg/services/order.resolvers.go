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
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrder) (*entity.Order, error) {
	panic(fmt.Errorf("not implemented: CreateOrder - createOrder"))
}

// UpdateOrder is the resolver for the updateOrder field.
func (r *mutationResolver) UpdateOrder(ctx context.Context, input model.UpdateOrderModel) (*entity.Order, error) {
	panic(fmt.Errorf("not implemented: UpdateOrder - updateOrder"))
}

// DeleteOrder is the resolver for the deleteOrder field.
func (r *mutationResolver) DeleteOrder(ctx context.Context, id int) (*entity.Order, error) {
	panic(fmt.Errorf("not implemented: DeleteOrder - deleteOrder"))
}

// Order is the resolver for the order field.
func (r *orderResolver) Order(ctx context.Context, obj *entity.Order) (*entity.Order, error) {
	return repositories.For(ctx).OrderLoader.Load(obj.ID)
}

// Orders is the resolver for the Orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*entity.Order, error) {
	panic(fmt.Errorf("not implemented: Orders - Orders"))
}

// OrderByAccountID is the resolver for the OrderByAccountId field.
func (r *queryResolver) OrderByAccountID(ctx context.Context) (*entity.Order, error) {
	panic(fmt.Errorf("not implemented: OrderByAccountID - OrderByAccountId"))
}

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

type orderResolver struct{ *Resolver }
