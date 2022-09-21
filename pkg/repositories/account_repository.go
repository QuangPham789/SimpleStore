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

func GetAllAccounts(ctx context.Context) ([]*entity.Accounts, error) {
	var lstAccountResult []*entity.Accounts
	err := dbmodels.NewQuery(
		qm.Select("accounts.id as id", "accounts.username as username", "accounts.password as password", "accounts.created_at as created_at"),
		qm.From(dbmodels.TableNames.Accounts),
	).Bind(ctx, boil.GetDB(), &lstAccountResult)

	if err != nil {
		log.Fatal("Get all account fail", err)
	}
	return lstAccountResult, nil
}

func GetAccountByID(ctx context.Context, id int64) (*entity.Accounts, error) {
	var accountResult = entity.Accounts{}

	err := dbmodels.Accounts(qm.Where("id=?", id)).Bind(ctx, boil.GetDB(), &accountResult)

	if err != nil {
		log.Fatal("Get account by ID fail", err)
	}

	return &accountResult, nil

}

func CreateAccount(ctx context.Context, input input.NewAccount) (*entity.Accounts, error) {
	var account = dbmodels.Account{
		Username: input.Username,
		Password: int64(input.Password),
	}

	err := account.Insert(context.Background(), boil.GetContextDB(), boil.Infer())
	if err != nil {
		log.Fatal("Insert account fail", err)
	}

	accountResult := entity.Accounts{
		ID:       int(account.ID),
		Username: account.Username,
		Password: int(account.Password),
	}

	return &accountResult, nil

}

func UpdateAccount(ctx context.Context, input input.UpdateAccountModel) (*entity.Accounts, error) {
	accountUpdate, err := dbmodels.FindAccount(ctx, boil.GetContextDB(), int64(input.ID))

	if err != nil {
		log.Fatal("Cannot find accounts with ID: ", input.ID, err)
	}

	accountUpdate.Username = input.Username
	accountUpdate.Password = int64(input.Password)

	rows, err := accountUpdate.Update(ctx, boil.GetContextDB(), boil.Infer())

	var accountResult entity.Accounts

	if rows > 0 {
		accountResult = entity.Accounts{
			Username: accountUpdate.Username,
			Password: int(accountUpdate.Password),
		}
	}
	return &accountResult, nil

}