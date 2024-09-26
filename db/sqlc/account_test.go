package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/Chimboys/simplebank/util"
	"github.com/stretchr/testify/require"
)

// Like the idea of creating a random account for testing purposes
// That allows us to create a new random account and for each test case
func createRandomAccount(t *testing.T) Account {

	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
func TestCreateAccount(t *testing.T) {
	account := createRandomAccount(t)
	account2, err := testQueries.GetAccountByID(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, account.Balance, account2.Balance)
	require.Equal(t, account.Currency, account2.Currency)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.CreatedAt, account2.CreatedAt)
	require.WithinDuration(t, account.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccountBalance(t *testing.T) {
	account := createRandomAccount(t)
	arg := UpdateAccountBalanceParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccountBalance(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, account2)
	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account.Currency, account2.Currency)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.CreatedAt, account2.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	deletedAccount, err := testQueries.DeleteAccountByOwner(context.Background(), account.Owner)
	require.NoError(t, err)
	require.NotEmpty(t, deletedAccount)

	// Verify that the deleted account matches the created account
	require.Equal(t, account.ID, deletedAccount.ID)
	require.Equal(t, account.Owner, deletedAccount.Owner)
	require.Equal(t, account.Balance, deletedAccount.Balance)
	require.Equal(t, account.Currency, deletedAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, deletedAccount.CreatedAt, 0)

	// Try to delete the same account again, it should return an error
	_, err = testQueries.DeleteAccountByOwner(context.Background(), account.Owner)
	require.Error(t, err)
}
