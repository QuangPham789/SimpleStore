package repositories

import (
	"context"
	"log"

	// "graphdemo/pkg/dataloader"

	dbmodels "graphdemo/pkg/db/models"
	"graphdemo/pkg/entity"
	input "graphdemo/pkg/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
func GetOrderByID(ctx context.Context, ids []int) ([]*entity.Order, error) {
	var orderResult []*entity.Order
	convertedIDs := make([]interface{}, len(ids))
	for index, num := range ids {
		convertedIDs[index] = num
	}
	err := dbmodels.Orders(qm.WhereIn("id IN ?", convertedIDs...)).Bind(ctx, boil.GetDB(), &orderResult)

	if err != nil {
		log.Fatal("Get item by id fail", err)
	}

	return orderResult, nil
}
func GetAllOrder(ctx context.Context) ([]*entity.Order, error) {
	var lstOrderResult []*entity.Order
	// err := dbmodels.NewQuery(
	// 	qm.Select("order.id as id", "order.code as code", "order.total_price as totalprice"),
	// 	qm.From(dbmodels.TableNames.Order),
	// ).Bind(ctx, boil.GetDB(), &lstOrderResult)

	// if err != nil {
	// 	log.Fatal("Get all orders fail: ", err)
	// }

	ordersFromDB, err := dbmodels.Orders().All(ctx, boil.GetContextDB())

	if err != nil {
		log.Fatal("Get all orders fail: ", err)
	}

	for _, o := range ordersFromDB {
		orderItem := entity.Order{
			ID:         int(o.ID),
			Code:       o.Code,
			Accountid:  int(o.AccountID),
			TotalPrice: int(o.TotalPrice),
		}
		lstOrderResult = append(lstOrderResult, &orderItem)
	}
	return lstOrderResult, nil
}

func CreateOrder(ctx context.Context, input input.NewOrder) (*entity.Order, error) {
	var order = dbmodels.Order{
		AccountID:  int64(input.Accountid),
		TotalPrice: int64(input.Total),
	}

	err := order.Insert(context.Background(), boil.GetContextDB(), boil.Infer())
	if err != nil {
		log.Println("Insert order fail", err)
	}

	orderResult := entity.Order{
		Accountid:  int(order.AccountID),
		TotalPrice: int(order.TotalPrice),
	}

	return &orderResult, nil

}
