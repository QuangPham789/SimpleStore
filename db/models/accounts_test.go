package models

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

func createRandomAccount(t *testing.T) Account {
	db := util.InitDB()
	defer db.Close()
	var account = Account{
		Username: util.RandomUsername(),
		Password: util.RandomPassword(),
	}

	err := account.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}
	return account

}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)

}

func TestGetAccountByID(t *testing.T) {
	db := util.InitDB()
	defer db.Close()
	accountTest := createRandomAccount(t)

	accountResult, err := Accounts(qm.Where("id=?", accountTest.ID)).One(context.Background(), db)

	require.NoError(t, err)
	require.NotEmpty(t, accountResult)
	require.Equal(t, accountTest.ID, accountResult.ID)
	require.Equal(t, accountTest.Username, accountResult.Username)
	require.Equal(t, accountTest.Password, accountResult.Password)
	require.WithinDuration(t, accountTest.CreatedAt, accountResult.CreatedAt, time.Second)

}

func findAccountByID(accountID int64) *Account {
	// init and close db connection when finish
	db := util.InitDB()
	defer db.Close()

	var account, err = FindAccount(context.Background(), db, accountID)
	if err != nil {
		log.Fatal("Can't not find account with id: ", accountID, "Error: ", err)
	}
	return account
}

func TestUpdateAccount(t *testing.T) {
	// init and close db connection when finish
	db := util.InitDB()
	defer db.Close()

	accountTest := createRandomAccount(t)

	accountUpdate := findAccountByID(accountTest.ID)
	accountUpdate.Username = "Quang"

	rowsAff, err := accountUpdate.Update(context.Background(), db, boil.Infer())

	require.NoError(t, err)
	require.NotEmpty(t, rowsAff)
	require.Equal(t, accountTest.ID, accountUpdate.ID)

}

func TestDeleteAccount(t *testing.T) {
	db := util.InitDB()
	defer db.Close()

	accountTest := createRandomAccount(t)

	accountDelete := findAccountByID(accountTest.ID)

	rowsAff, err := accountDelete.Delete(context.Background(), db)
	require.NoError(t, err)
	require.NotEmpty(t, rowsAff)

}
