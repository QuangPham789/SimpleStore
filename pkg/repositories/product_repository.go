package repositories

import (
	"context"
	dbmodels "graphdemo/pkg/db/models"
	entity "graphdemo/pkg/entity"
	input "graphdemo/pkg/graph/model"
	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetAllProduct(ctx context.Context) ([]*entity.Product, error) {
	var lstProductResult []*entity.Product
	err := dbmodels.NewQuery(
		qm.Select("products.id as id", "products.code as code", "products.name as name", "products.description as description", "products.category as category", "products.price as price"),
		qm.From(dbmodels.TableNames.Products),
	).Bind(ctx, boil.GetDB(), &lstProductResult)

	if err != nil {
		log.Fatal("Get all products fail", err)
	}
	return lstProductResult, nil
}

func GetProductByCode(ctx context.Context, code string) (*entity.Product, error) {
	var productResult = entity.Product{}

	err := dbmodels.Products(qm.Where("code=?", code)).Bind(ctx, boil.GetDB(), &productResult)

	if err != nil {
		log.Fatal("Get account by code fail", err)
	}

	return &productResult, nil

}

func CreateProduct(ctx context.Context, input input.NewProduct) (*entity.Product, error) {
	var product = dbmodels.Product{
		Code:        input.Code,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Category:    input.Category,
	}

	err := product.Insert(context.Background(), boil.GetContextDB(), boil.Infer())
	if err != nil {
		log.Println("Insert product fail", err)
	}

	productResult := entity.Product{
		Code:        product.Code,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
	}

	return &productResult, nil

}

func UpdateProduct(ctx context.Context, input input.UpdateProductModel) (*entity.Product, error) {
	productUpdate, err := dbmodels.FindProduct(ctx, boil.GetContextDB(), int64(input.ID))

	if err != nil {
		log.Fatal("Cannot find product with ID: ", input.ID, err)
	}

	productUpdate.Code = input.Code
	productUpdate.Name = input.Name
	productUpdate.Description = input.Description
	productUpdate.Price = input.Price
	productUpdate.Category = input.Category

	rows, err := productUpdate.Update(ctx, boil.GetContextDB(), boil.Infer())

	var productResult entity.Product

	if rows > 0 {
		productResult = entity.Product{
			Code:        productUpdate.Code,
			Name:        productUpdate.Name,
			Description: productUpdate.Description,
			Price:       productUpdate.Price,
			Category:    productUpdate.Category,
		}
	}
	return &productResult, nil

}
