package main

import (
	"context"

	"gorm.io/gorm"

	"github.com/bar-raisers/intention/common/resources/db"
	accounts "github.com/bar-raisers/intention/finance/contracts/accounts/v1"
)

type AccountsService struct {
	db *gorm.DB
}

func NewAccountsService() *AccountsService {
	dataSource := db.NewDataSource(
		"localhost",
		"postgres",
		"postgres",
		"db",
		5432,
		"America/Vancouver",
	)
	db := db.Connect(dataSource)

	return &AccountsService{
		db: db,
	}
}

func (service *AccountsService) CreateTransaction(ctx context.Context, request *accounts.CreateTransactionRequest) (*accounts.CreateTransactionResponse, error) {
	return &accounts.CreateTransactionResponse{}, nil
}

func (service *AccountsService) ListTransactions(ctx context.Context, request *accounts.ListTransactionsRequest) (*accounts.ListTransactionsResponse, error) {
	return &accounts.ListTransactionsResponse{}, nil
}
