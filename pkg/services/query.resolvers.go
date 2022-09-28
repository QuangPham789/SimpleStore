package services

import (
	"context"
	"fmt"
	"graphdemo/pkg/entity"
	"graphdemo/pkg/graph/generated"
	"graphdemo/pkg/repositories"
	"sync"
)

// account
// Accounts is the resolver for the Accounts field.
func (r *queryResolver) Accounts(ctx context.Context) ([]*entity.Accounts, error) {
	accountsFromDB, err := repositories.GetAllAccounts(ctx)
	if err != nil {
		return nil, err
	}

	return accountsFromDB, nil
}

// AccountByID is the resolver for the AccountByID field.
func (r *queryResolver) AccountByID(ctx context.Context, id int) (*entity.Accounts, error) {
	// accountsFromDB, err := repositories.GetAccountByID(ctx, id)
	// if err != nil {
	// 	return nil, err
	// }

	// return accountsFromDB, nil
	workers := 3
	total := 5
	jobs := make(chan int)
	results := make(chan *entity.Accounts)
	wg := &sync.WaitGroup{}
	// start workers
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go repositories.Worker(jobs, results, ctx, id, wg)
	}

	// insert jobs in background
	go func() {
		for j := 0; j < total; j++ {
			jobs <- j
		}
		close(jobs)
		wg.Wait()
		// all workers are done so no more results
		close(results)
	}()
	var accountsFromDB *entity.Accounts
	// collect results
	for i := 0; i < total; i++ {
		accountsFromDB = <-results
	}
	return accountsFromDB, nil
}

// Items is the resolver for the Items field.
func (r *queryResolver) Items(ctx context.Context) ([]*entity.Item, error) {
	itemFromDB, err := repositories.GetAllItem(ctx)
	if err != nil {
		return nil, err
	}

	return itemFromDB, nil
}

// ItemByOrderID is the resolver for the ItemByOrderID field.
func (r *queryResolver) ItemByOrderID(ctx context.Context, id int) ([]*entity.Item, error) {
	panic(fmt.Errorf("not implemented: ItemByOrderID - ItemByOrderID"))
}

// Orders is the resolver for the Orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*entity.Order, error) {
	ordersFromDB, err := repositories.GetAllOrder(ctx)
	if err != nil {
		return nil, err
	}

	return ordersFromDB, nil
}

// OrderByAccountID is the resolver for the OrderByAccountId field.
func (r *queryResolver) OrderByAccountID(ctx context.Context) (*entity.Order, error) {
	panic(fmt.Errorf("not implemented: OrderByAccountID - OrderByAccountId"))
}

// Products is the resolver for the Products field.
func (r *queryResolver) Products(ctx context.Context) ([]*entity.Product, error) {
	productFromDB, err := repositories.GetAllProduct(ctx)
	if err != nil {
		return nil, err
	}

	return productFromDB, nil
}

// ProductByCode is the resolver for the ProductByCode field.
func (r *queryResolver) ProductByCode(ctx context.Context, code string) (*entity.Product, error) {
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

// GenToken is the resolver for the GenToken field.

type queryResolver struct{ *Resolver }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }
