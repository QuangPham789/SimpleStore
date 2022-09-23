// package dataloader
package repositories

import (
	"context"
	"graphdemo/pkg/entity"
	"log"
	"net/http"
	"time"
)

const (
	// accountLoaderKey = "accountloader"
	itemLoaderKey = "itemloader"
)

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

func DataloaderMiddleware(ctx context.Context, next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemLoader := ItemLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(keys []string) ([]*entity.Item, []error) {
				itemsFromDB, err := GetItemByOrderID(ctx, keys)
				log.Println("loader", len(itemsFromDB))
				return itemsFromDB, []error{err}
			},
		}

		ctx := context.WithValue(r.Context(), itemLoaderKey, &itemLoader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetItemLoader(ctx context.Context) *ItemLoader {
	return ctx.Value(itemLoaderKey).(*ItemLoader)
}
