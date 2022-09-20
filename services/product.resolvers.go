package services

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphdemo/graph/model"
	"graphdemo/models"
	"graphdemo/repositories"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*models.Product, error) {
	panic(fmt.Errorf("not implemented: CreateProduct - createProduct"))
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, input model.UpdateProductModel) (*models.Product, error) {
	panic(fmt.Errorf("not implemented: UpdateProduct - updateProduct"))
}

// GetAllProduct is the resolver for the GetAllProduct field.
func (r *queryResolver) GetAllProduct(ctx context.Context) ([]*models.Product, error) {
	productFromDB, err := repositories.GetAllProduct(ctx)
	if err != nil {
		return nil, err
	}

	return productFromDB, nil
}

// GetProductByCode is the resolver for the GetProductByCode field.
func (r *queryResolver) GetProductByCode(ctx context.Context, code string) (*models.Product, error) {
	panic(fmt.Errorf("not implemented: GetProductByCode - GetProductByCode"))
}
