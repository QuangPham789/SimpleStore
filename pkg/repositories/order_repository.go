package repositories

import (
	"context"
	dbmodels "graphdemo/pkg/db/models"
	"graphdemo/pkg/entity"
	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetItemsInOrderID(ctx context.Context, code int64) ([]*entity.Item, error) {
	var itemResult []*entity.Item

	// err := dbmodels.Items(qm.Where("code=?", code)).Bind(ctx, boil.GetDB(), &itemResult)
	err := dbmodels.Orders(
		dbmodels.OrderWhere.ID.EQ(code),
		qm.Load(dbmodels.OrderRels.ItemID)).Bind(ctx, boil.GetDB(), &itemResult)
	if err != nil {
		log.Fatal("Get item by code fail", err)
	}

	return itemResult, nil

}
