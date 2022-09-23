package services

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphdemo/pkg/entity"
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

// GetAllOrder is the resolver for the GetAllOrder field.
func (r *queryResolver) GetAllOrder(ctx context.Context) ([]*entity.Order, error) {
	orderFromDB, err := repositories.GetAllOrder(ctx)
	if err != nil {
		return nil, err
	}

	return orderFromDB, nil
}

// GetOrderByAccountID is the resolver for the GetOrderByAccountId field.
func (r *queryResolver) GetOrderByAccountID(ctx context.Context) (*entity.Order, error) {
	panic(fmt.Errorf("not implemented: GetOrderByAccountID - GetOrderByAccountId"))
}
