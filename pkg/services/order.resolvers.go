package services

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphdemo/pkg/graph/model"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrder) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: CreateOrder - createOrder"))
}

// UpdateOrder is the resolver for the updateOrder field.
func (r *mutationResolver) UpdateOrder(ctx context.Context, input model.UpdateOrderModel) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: UpdateOrder - updateOrder"))
}

// GetAllOrder is the resolver for the GetAllOrder field.
func (r *queryResolver) GetAllOrder(ctx context.Context) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented: GetAllOrder - GetAllOrder"))
}

// GetOrderByAccountID is the resolver for the GetOrderByAccountId field.
func (r *queryResolver) GetOrderByAccountID(ctx context.Context) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: GetOrderByAccountID - GetOrderByAccountId"))
}
