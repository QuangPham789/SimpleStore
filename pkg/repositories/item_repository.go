package repositories

import (
	"context"
	dbmodels "graphdemo/pkg/db/models"
	"graphdemo/pkg/entity"
	input "graphdemo/pkg/graph/model"
	"log"
	"strconv"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetAllItem(ctx context.Context) ([]*entity.Item, error) {
	var lstItemResult []*entity.Item
	err := dbmodels.NewQuery(
		qm.Select("item.id as id", "item.code as code", "item.product_id as product_id", "item.unit as unit", "item.total_price as total_price"),
		qm.From(dbmodels.TableNames.Item),
	).Bind(ctx, boil.GetDB(), &lstItemResult)

	if err != nil {
		log.Fatal("Get all products fail", err)
	}
	return lstItemResult, nil
}

func CreateItem(ctx context.Context, input input.NewItem) (*entity.Item, error) {
	//parse price
	totalPrice, err := strconv.ParseInt(input.Total, 10, 0)
	if err != nil {
		log.Println("CreateItem", "parse price fail", err)
	}
	var item = dbmodels.Item{
		Code:       input.Code,
		ProductID:  input.Productid,
		Unit:       input.Unit,
		TotalPrice: totalPrice,
	}

	err = item.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		log.Println("Insert item fail", err)
	}

	itemResult := entity.Item{
		Code:      item.Code,
		Productid: item.ProductID,
		Unit:      item.Unit,
		Price:     input.Price,
		Total:     strconv.Itoa(int(totalPrice)),
	}

	return &itemResult, nil

}

func UpdateItem(ctx context.Context, input input.ItemModel) (*entity.Item, error) {
	itemUpdate, err := dbmodels.FindItem(ctx, boil.GetContextDB(), int64(input.ID))
	//parse price
	totalPrice, err := strconv.ParseInt(input.Total, 10, 0)
	if err != nil {
		log.Println("CreateItem", "parse price fail", err)
	}
	if err != nil {
		log.Fatal("Cannot find product with ID: ", input.ID, err)
	}

	itemUpdate.Code = input.Code
	itemUpdate.ProductID = input.Productid
	itemUpdate.Unit = input.Unit
	itemUpdate.TotalPrice = totalPrice

	rows, err := itemUpdate.Update(ctx, boil.GetContextDB(), boil.Infer())

	var itemResult entity.Item

	if rows > 0 {
		itemResult = entity.Item{
			Code:      itemUpdate.Code,
			Unit:      itemUpdate.Unit,
			Productid: itemUpdate.ProductID,
			Price:     input.Price,
			Total:     input.Total,
		}
	}
	return &itemResult, nil

}

func GetItemByOrderID(ctx context.Context, code string) (*entity.Item, error) {
	var itemResult = entity.Item{}

	err := dbmodels.Items(qm.Where("code=?", code)).Bind(ctx, boil.GetDB(), &itemResult)

	if err != nil {
		log.Fatal("Get item by code fail", err)
	}

	return &itemResult, nil

}
