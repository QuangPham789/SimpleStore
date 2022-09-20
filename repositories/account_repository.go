package repositories

import (
	"context"
	dbmodels "graphdemo/db/models"
	input "graphdemo/graph/model"
	models "graphdemo/models"
	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetAllAccounts(ctx context.Context) ([]*models.Accounts, error) {
	var lstAccountResult []*models.Accounts
	err := dbmodels.NewQuery(
		qm.Select("accounts.id as id", "accounts.username as username", "accounts.password as password", "accounts.created_at as created_at"),
		qm.From(dbmodels.TableNames.Accounts),
	).Bind(ctx, boil.GetDB(), &lstAccountResult)

	if err != nil {
		log.Fatal("Get all account fail", err)
	}
	return lstAccountResult, nil
}

func GetAccountByID(ctx context.Context, id int64) (*models.Accounts, error) {
	var accountResult = models.Accounts{}

	err := dbmodels.Accounts(qm.Where("id=?", id)).Bind(ctx, boil.GetDB(), &accountResult)

	if err != nil {
		log.Fatal("Get account by ID fail", err)
	}

	return &accountResult, nil

}

func CreateAccount(ctx context.Context, input input.NewAccount) (*models.Accounts, error) {
	var account = dbmodels.Account{
		Username: input.Username,
		Password: int64(input.Password),
	}

	err := account.Insert(context.Background(), boil.GetContextDB(), boil.Infer())
	if err != nil {
		log.Fatal("Insert account fail", err)
	}

	accountResult := models.Accounts{
		ID:       int(account.ID),
		Username: account.Username,
		Password: int(account.Password),
	}

	return &accountResult, nil

}

func UpdateAccount(ctx context.Context, input input.UpdateAccountModel) (*models.Accounts, error) {
	accountUpdate, err := dbmodels.FindAccount(ctx, boil.GetContextDB(), int64(input.ID))

	if err != nil {
		log.Fatal("Cannot find accounts with ID: ", input.ID, err)
	}

	accountUpdate.Username = input.Username
	accountUpdate.Password = int64(input.Password)

	rows, err := accountUpdate.Update(ctx, boil.GetContextDB(), boil.Infer())

	var accountResult models.Accounts

	if rows > 0 {
		accountResult = models.Accounts{
			Username: accountUpdate.Username,
			Password: int(accountUpdate.Password),
		}
	}
	return &accountResult, nil

}

func GetAllProduct(ctx context.Context) ([]*models.Product, error) {
	var lstProductResult []*models.Product
	err := dbmodels.NewQuery(
		qm.Select("products.id as id", "products.code as code", "products.name as name", "products.description as description", "products.category as category", "products.price as price"),
		qm.From(dbmodels.TableNames.Products),
	).Bind(ctx, boil.GetDB(), &lstProductResult)

	if err != nil {
		log.Fatal("Get all products fail", err)
	}
	return lstProductResult, nil
}
