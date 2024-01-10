package main

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/bar-raisers/intention/common/resources/db"

	contract_accounts "github.com/bar-raisers/intention/finance/contracts/accounts/v1"
	"github.com/bar-raisers/intention/finance/services/accounts/models"
	"github.com/bar-raisers/intention/finance/services/accounts/resources"
	"github.com/bar-raisers/intention/finance/services/accounts/utilities"
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
	resources.Migrate(db)

	return &AccountsService{
		db: db,
	}
}

func (service *AccountsService) CreateTransaction(
	ctx context.Context,
	request *contract_accounts.CreateTransactionRequest,
) (*contract_accounts.CreateTransactionResponse, error) {
	transaction_pb := request.GetTransaction()
	if transaction_pb == nil {
		return nil, fmt.Errorf("no transaction provided")
	}

	transaction := models.NewTransactionFromProto(transaction_pb)
	err := resources.CreateTransaction(service.db, transaction)
	if err != nil {
		return nil, err
	}

	return &contract_accounts.CreateTransactionResponse{}, nil
}

func (service *AccountsService) ListTransactions(
	ctx context.Context,
	request *contract_accounts.ListTransactionsRequest,
) (*contract_accounts.ListTransactionsResponse, error) {
	transactions := resources.ListTransactions(service.db)
	transactions_pb := utilities.ConvertTransactionsToProto(transactions)

	response := &contract_accounts.ListTransactionsResponse{
		Transactions: transactions_pb,
	}
	return response, nil
}
