package repositories

import (
	"context"
	"fmt"
	dbmodels "graphdemo/pkg/db/models"
	entity "graphdemo/pkg/entity"
	input "graphdemo/pkg/graph/model"
	"log"
	"sync"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func Worker(jobs <-chan int, results chan<- *entity.Accounts, ctx context.Context, id int, wg *sync.WaitGroup) {
	for _ = range jobs {
		account, err := GetAccountByID(ctx, id)
		if err != nil {
			log.Println(err)

		}
		results <- account
	}
	wg.Done()
}

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

func GetAccountByID(ctx context.Context, id int) (*entity.Accounts, error) {
	// var accountResult = entity.Accounts{}

	// err := dbmodels.Accounts(qm.Where("id=?", id)).Bind(ctx, boil.GetDB(), &accountResult)

	// if err != nil {
	// 	log.Fatal("Get account by ID fail", err)
	// }

	// return &accountResult, nil
	return GetAccountByField("id", id, ctx)

}

func CreateAccount(ctx context.Context, input input.NewAccount) (*entity.Accounts, error) {
	// passwordHash, err := util.HashPassword(input.Password)
	// if err != nil {
	// 	log.Fatal("Hash password account fail", err)
	// }
	var account = dbmodels.Account{
		Username: input.Username,
		Password: int64(input.Password),
	}

	err := account.Insert(context.Background(), boil.GetContextDB(), boil.Infer())
	if err != nil {
		log.Fatal("Insert account fail", err)
	}

	// token, err := GenToken()
	accountResult := entity.Accounts{
		ID:       int(account.ID),
		Username: account.Username,
		Password: int(account.Password),
		// Token:     token.AccessToken,
		// ExpiredAt: token.ExpiredAt.String(),
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
			// Password: accountUpdate.Password,
		}
	}
	return &accountResult, nil

}

func DeleteAccount(ctx context.Context, id int) (*entity.Accounts, error) {
	accountDelete, err := dbmodels.FindAccount(ctx, boil.GetContextDB(), int64(id))
	if err != nil {
		log.Fatal("Cannot find product with ID: ", id, err)
	}
	rows, err := accountDelete.Delete(ctx, boil.GetContextDB())

	var accountResult entity.Accounts
	if rows > 0 {
		accountResult = entity.Accounts{
			ID:       id,
			Username: accountDelete.Username,
			// Password: accountDelete.Password,
		}
	}

	return &accountResult, nil
}

func GetAccountByField(field string, value int, ctx context.Context) (*entity.Accounts, error) {
	var accountResult = entity.Accounts{}

	err := dbmodels.Accounts(qm.Where(fmt.Sprintf("%v = ?", field), value)).Bind(ctx, boil.GetDB(), &accountResult)

	return &accountResult, err
}
