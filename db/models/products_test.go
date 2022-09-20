package dbmodels

import (
	"context"
	"graphdemo/util"
	"log"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func createRandomProduct(t *testing.T) Product {
	db := util.InitDB()
	defer db.Close()
	var product = Product{
		Code:        util.RandomProductCode(),
		Name:        util.RandomProductName(),
		Description: util.RandomProductName(),
		Category:    util.RandomProductName(),
		Price:       util.RandomProductPrice(),
	}

	err := product.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}
	return product

}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)

}

func TestGetProductByCode(t *testing.T) {
	db := util.InitDB()
	defer db.Close()
	productTest := createRandomProduct(t)

	productResult, err := Products(qm.Where("code=?", productTest.Code)).One(context.Background(), db)

	require.NoError(t, err)
	require.NotEmpty(t, productResult)
	require.Equal(t, productTest.ID, productResult.ID)
	require.Equal(t, productTest.Code, productResult.Code)
	require.Equal(t, productTest.Name, productResult.Name)
	require.WithinDuration(t, productTest.CreatedAt, productResult.CreatedAt, time.Second)

}

func findProductByID(productID int64) *Product {
	// init and close db connection when finish
	db := util.InitDB()
	defer db.Close()

	var product, err = FindProduct(context.Background(), db, productID)
	if err != nil {
		log.Fatal("Can't not find product with id: ", productID, "Error: ", err)
	}
	return product
}

func TestUpdateProduct(t *testing.T) {
	// init and close db connection when finish
	db := util.InitDB()
	defer db.Close()

	productTest := createRandomProduct(t)

	productUpdate := findProductByID(productTest.ID)

	productUpdate.Name = "IP14"

	rowsAff, err := productUpdate.Update(context.Background(), db, boil.Infer())

	require.NoError(t, err)
	require.NotEmpty(t, rowsAff)
	require.Equal(t, productTest.ID, productUpdate.ID)

}

func TestDeleteProduct(t *testing.T) {
	db := util.InitDB()
	defer db.Close()

	productTest := createRandomProduct(t)

	productDelete := findProductByID(productTest.ID)

	rowsAff, err := productDelete.Delete(context.Background(), db)
	require.NoError(t, err)
	require.NotEmpty(t, rowsAff)

}
