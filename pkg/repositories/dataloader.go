// package dataloader
package repositories

import (
	"context"
	"graphdemo/pkg/entity"
	"log"
	"net/http"
	"time"
)

// const (
// 	// accountLoaderKey = "accountloader"
// 	// itemLoaderKey    = "itemloader"
// 	orderLoaderKey = "orderloader"
// )

type loadersString string

const loadersKey = loadersString("dataloaders")

type Loaders struct {
	OrderLoader OrderLoader
}

// func DataloaderMiddleware(ctx context.Context, next http.Handler) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		accountLoader := AccountsLoader{
// 			maxBatch: 100,
// 			wait:     1 * time.Millisecond,
// 			fetch: func(keys []int) ([]*entity.Accounts, []error) {
// 				accountsFromDB, err := GetAccountByID(ctx, keys)
// 				return accountsFromDB, []error{err}
// 			},
// 		}

// 		ctx := context.WithValue(r.Context(), accountLoaderKey, &accountLoader)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// func GetAccountLoader(ctx context.Context) *AccountsLoader {
// 	return ctx.Value(accountLoaderKey).(*AccountsLoader)
// }

// func DataItemloaderMiddleware(ctx context.Context, next http.Handler) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		itemLoader := ItemLoader{
// 			maxBatch: 100,
// 			wait:     1 * time.Millisecond,
// 			fetch: func(keys []string) ([]*entity.Item, []error) {
// 				itemsFromDB, err := GetItemByOrderID(ctx, keys)
// 				log.Println("loader", len(itemsFromDB))
// 				return itemsFromDB, []error{err}
// 			},
// 		}

// 		ctx := context.WithValue(r.Context(), itemLoaderKey, &itemLoader)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// func GetItemLoader(ctx context.Context) *ItemLoader {
// 	return ctx.Value(itemLoaderKey).(*ItemLoader)
// }

func DataloaderMiddleware(ctx context.Context, next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			OrderLoader: OrderLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(keys []int) ([]*entity.Order, []error) {
					orderFromDB, err := GetOrderByID(ctx, keys)
					if err != nil {
						log.Fatal(err)
					}

					order := make(map[int]*entity.Order, len(orderFromDB))

					for _, o := range orderFromDB {
						order[o.ID] = o
					}
					return orderFromDB, []error{err}
				},
			},
		})

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

//	func GetOrderLoader(ctx context.Context) *OrderLoader {
//		return ctx.Value(orderLoaderKey).(*OrderLoader)
//	}
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
