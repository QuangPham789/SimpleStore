package services

import (
	"context"
	"fmt"
	"graphdemo/pkg/entity"
	"graphdemo/pkg/graph/generated"
	"graphdemo/pkg/graph/model"
	"graphdemo/pkg/repositories"
	"log"
	"strconv"
)

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, input model.NewAccount) (*entity.Accounts, error) {
	accountResult, err := repositories.CreateAccount(ctx, input)
	if err != nil {
		return nil, err
	}
	return accountResult, nil
}

// UpdateAccount is the resolver for the updateAccount field.
func (r *mutationResolver) UpdateAccount(ctx context.Context, input model.UpdateAccountModel) (*entity.Accounts, error) {
	accountResult, err := repositories.UpdateAccount(ctx, input)
	if err != nil {
		return nil, err
	}
	return accountResult, nil
}

// DeleteAccount is the resolver for the deleteAccount field.
func (r *mutationResolver) DeleteAccount(ctx context.Context, id int) (*entity.Accounts, error) {
	accountResult, err := repositories.DeleteAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	return accountResult, nil
}

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
	totalPriceStr := strconv.Itoa(int(totalPrice))
	input.Total = totalPriceStr
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
