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

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*entity.Product, error) {
	productResult, err := repositories.CreateProduct(ctx, input)
	if err != nil {
		return nil, err
	}
	return productResult, nil
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, input model.UpdateProductModel) (*entity.Product, error) {
	productResult, err := repositories.UpdateProduct(ctx, input)
	if err != nil {
		return nil, err
	}
	return productResult, nil
}

// GetAllProduct is the resolver for the GetAllProduct field.
func (r *queryResolver) GetAllProduct(ctx context.Context) ([]*entity.Product, error) {
	productFromDB, err := repositories.GetAllProduct(ctx)
	if err != nil {
		return nil, err
	}

	return productFromDB, nil
}

// GetProductByCode is the resolver for the GetProductByCode field.
func (r *queryResolver) GetProductByCode(ctx context.Context, code string) (*entity.Product, error) {
	panic(fmt.Errorf("not implemented: GetProductByCode - GetProductByCode"))
}

type newProductResolver struct{ *Resolver }
type updateProductModelResolver struct{ *Resolver }
