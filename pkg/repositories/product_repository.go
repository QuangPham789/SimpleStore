package repositories

import (
	"context"
	dbmodels "graphdemo/pkg/db/models"
	entity "graphdemo/pkg/entity"
	input "graphdemo/pkg/graph/model"
	"log"
	"os"

	// "github.com/graph-gophers/dataloader"
	// "github.com/graph-gophers/dataloader"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// UserReader reads Users from a database
// type ProductReader struct {
// 	conn *sql.DB
// }

// // dataloader
// func GetProducts(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
// 	// read all requested users in a single query
// 	productIDs := make([]string, len(keys))
// 	for ix, key := range keys {
// 		productIDs[ix] = key.String()
// 		log.Println(productIDs[ix])
// 	}
// 	products, err := dbmodels.Products(
// 		qm.WhereIn("id=?", productIDs),
// 	).All(ctx, boil.GetContextDB())

// 	if err != nil {
// 		return nil
// 	}

// 	// return User records into a map by ID
// 	productById := map[string]*dbmodels.Product{}

// 	for _, item := range products {
// 		productId := strconv.Itoa(int(item.ID))
// 		productById[productId] = item
// 	}

// 	// return users in the same order requested
// 	output := make([]*dataloader.Result, len(keys))
// 	for index, userKey := range keys {
// 		product, ok := productById[userKey.String()]
// 		if ok {
// 			output[index] = &dataloader.Result{Data: product, Error: nil}
// 		} else {
// 			err := fmt.Errorf("user not found %s", userKey.String())
// 			output[index] = &dataloader.Result{Data: nil, Error: err}
// 		}
// 	}
// 	return output
// }

// // Loaders wrap your data loaders to inject via middleware
// type Loaders struct {
// 	ProductLoader *dataloader.Loader
// }

// // NewLoaders instantiates data loaders for the middleware
// func NewLoaders(conn *sql.DB) *Loaders {
// 	// define the data loader
// 	productReader := boil.GetContextDB()
// 	loaders := &Loaders{
// 		ProductLoader: dataloader.NewBatchedLoader(productReader.),
// 	}
// 	return loaders
// }
//dataloader

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

	// err := dbmodels.Products(qm.WhereIn("code in ?", code)).Bind(ctx, boil.GetDB(), &productResult)
	err := dbmodels.Products(qm.WhereIn("code in ?", code)).Bind(ctx, boil.GetContextDB(), &productResult)

	if err != nil {
		log.Fatal("Get account by code fail", err)
	}

	return &productResult, nil

}

func GetProductByID(ctx context.Context, id int) (*entity.Product, error) {

	boil.DebugMode = true

	// Optionally set the writer as well. Defaults to os.Stdout
	fh, _ := os.Open("debug.txt")
	boil.DebugWriter = fh
	var productResult = entity.Product{}

	err := dbmodels.Products(qm.WhereIn("id in ?", id)).Bind(ctx, boil.GetDB(), &productResult)

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

func DeleteProduct(ctx context.Context, id int) (*entity.Product, error) {
	productDelete, err := dbmodels.FindProduct(ctx, boil.GetContextDB(), int64(id))
	if err != nil {
		log.Fatal("Cannot find product with ID: ", id, err)
	}
	rows, err := productDelete.Delete(ctx, boil.GetContextDB())

	var productResult entity.Product
	if rows > 0 {
		productResult = entity.Product{
			ID:          id,
			Code:        productDelete.Code,
			Name:        productDelete.Name,
			Description: productDelete.Description,
			Price:       productDelete.Price,
			Category:    productDelete.Category,
		}
	}

	return &productResult, nil
}
