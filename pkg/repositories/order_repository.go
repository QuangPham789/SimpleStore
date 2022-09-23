package repositories

import (
	"context"
	// "graphdemo/pkg/dataloader"
	dbmodels "graphdemo/pkg/db/models"
	"graphdemo/pkg/entity"
	"log"
	"strconv"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// func GetItemsInOrderID(ctx context.Context, code int64) ([]*entity.Item, error) {
// 	var itemResult []*entity.Item

// 	// err := dbmodels.Items(qm.Where("code=?", code)).Bind(ctx, boil.GetDB(), &itemResult)
// 	err := dbmodels.Orders(
// 		dbmodels.OrderWhere.ID.EQ(code),
// 		qm.Load(dbmodels.OrderRels.ItemID)).Bind(ctx, boil.GetDB(), &itemResult)
// 	if err != nil {
// 		log.Fatal("Get item by code fail", err)
// 	}

// 	return itemResult, nil

// }

func GetAllOrder(ctx context.Context) ([]*entity.Order, error) {
	var lstOrderResult []*entity.Order

	var lstItem []*entity.Item
	lstOderFormDB, err := dbmodels.Orders().All(ctx, boil.GetContextDB())

	if err != nil {
		log.Fatal(err)
	}

	log.Println(len(lstOderFormDB))

	for _, itemOrder := range lstOderFormDB {
		log.Println(itemOrder.Items)
		// parse to str
		itemId, _ := strconv.ParseInt(itemOrder.Items, 10, 0)
		item, err := GetItemLoader(ctx).Load(itemOrder.Items)
		if err != nil {
			log.Fatal(err)
		}
		lstItem = append(lstItem, item)
		itemOrder := entity.Order{
			ID:        strconv.Itoa(int(itemOrder.ID)),
			Accountid: int(itemOrder.AccountID),
			Total:     int(itemOrder.TotalPrice),
			Itemid:    int(itemId),
			Items:     lstItem,
		}
		lstOrderResult = append(lstOrderResult, &itemOrder)
	}

	return lstOrderResult, nil
}
