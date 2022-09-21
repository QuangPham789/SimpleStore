package services

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphdemo/pkg/entity"
	"graphdemo/pkg/graph/generated"
	"graphdemo/pkg/graph/model"
)

// CreateItem is the resolver for the createItem field.
func (r *mutationResolver) CreateItem(ctx context.Context, input model.NewItem) (*entity.Item, error) {
	panic(fmt.Errorf("not implemented: CreateItem - createItem"))
}

// UpdateItem is the resolver for the updateItem field.
func (r *mutationResolver) UpdateItem(ctx context.Context, input model.ItemModel) (*entity.Item, error) {
	panic(fmt.Errorf("not implemented: UpdateItem - updateItem"))
}

// GetItemByOrderID is the resolver for the GetItemByOrderID field.
func (r *queryResolver) GetItemByOrderID(ctx context.Context) ([]*entity.Item, error) {
	panic(fmt.Errorf("not implemented: GetItemByOrderID - GetItemByOrderID"))
}

// ID is the resolver for the id field.
func (r *itemModelResolver) ID(ctx context.Context, obj *model.ItemModel, data int) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// ItemModel returns generated.ItemModelResolver implementation.
func (r *Resolver) ItemModel() generated.ItemModelResolver { return &itemModelResolver{r} }

type itemModelResolver struct{ *Resolver }
