package services

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
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

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, id int) (*entity.Product, error) {
	productResult, err := repositories.DeleteProduct(ctx, id)
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
	productFromDB, err := repositories.GetProductByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	return productFromDB, nil
	// product := entity.Product{
	// 	ID:          1,
	// 	Code:        "123",
	// 	Name:        "demo",
	// 	Description: "demo",
	// 	Price:       "1000",
	// 	Category:    "demo",
	// }
	// return &product, nil
}

type newProductResolver struct{ *Resolver }
type updateProductModelResolver struct{ *Resolver }
