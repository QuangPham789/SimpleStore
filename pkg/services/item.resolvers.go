package services

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphdemo/pkg/entity"
	"graphdemo/pkg/graph/model"
	"graphdemo/pkg/repositories"
	"log"
	"strconv"
)

// CreateItem is the resolver for the createItem field.
func (r *mutationResolver) CreateItem(ctx context.Context, input model.NewItem) (*entity.Item, error) {
	//get product by product id to get price
	product, err := repositories.GetProductByID(ctx, input.Productid)
	if err != nil {
		log.Println("CreateItem", "get product fail", err)
	}

	//parse price
	price, err := strconv.ParseInt(product.Price, 10, 0)
	if err != nil {
		log.Println("CreateItem", "parse price fail", err)
	}
	//calculate total price
	totalPrice := (price * int64(input.Unit))

	input.Total = strconv.Itoa(int(totalPrice))
	input.Price = strconv.Itoa(int(price))

	itemResult, err := repositories.CreateItem(ctx, input)
	if err != nil {
		return nil, err
	}
	return itemResult, nil
}

// UpdateItem is the resolver for the updateItem field.
func (r *mutationResolver) UpdateItem(ctx context.Context, input model.ItemModel) (*entity.Item, error) {
	//get product by product id to get price
	product, err := repositories.GetProductByID(ctx, input.Productid)
	if err != nil {
		log.Println("UpdateItem", "get product fail", err)
	}

	//parse price
	price, err := strconv.ParseInt(product.Price, 10, 0)
	if err != nil {
		log.Println("UpdateItem", "parse price fail", err)
	}
	//calculate total price
	totalPrice := (price * int64(input.Unit))

	input.Total = strconv.Itoa(int(totalPrice))
	input.Price = strconv.Itoa(int(price))
	log.Println("UpdateItem", "input.totalPrice", totalPrice)
	log.Println("UpdateItem", "input.Price", price)

	itemResult, err := repositories.UpdateItem(ctx, input)
	if err != nil {
		return nil, err
	}
	return itemResult, nil
}

// DeleteItem is the resolver for the deleteItem field.
func (r *mutationResolver) DeleteItem(ctx context.Context, id int) (*entity.Item, error) {
	panic(fmt.Errorf("not implemented: DeleteItem - deleteItem"))
}

// GetAllItem is the resolver for the GetAllItem field.
func (r *queryResolver) GetAllItem(ctx context.Context) ([]*entity.Item, error) {
	itemFromDB, err := repositories.GetAllItem(ctx)
	if err != nil {
		return nil, err
	}

	return itemFromDB, nil
}

// GetItemByOrderID is the resolver for the GetItemByOrderID field.
func (r *queryResolver) GetItemByOrderID(ctx context.Context, id int) ([]*entity.Item, error) {
	panic(fmt.Errorf("not implemented: GetItemByOrderID - GetItemByOrderID"))
}

type itemModelResolver struct{ *Resolver }
