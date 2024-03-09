package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/bar-raisers/intention/common/resources/db"

	contract_accounts "github.com/bar-raisers/intention/contracts/finance/accounts/v1"
	"github.com/bar-raisers/intention/services/finance/accounts/models"
	"github.com/bar-raisers/intention/services/finance/accounts/resources"
	"github.com/bar-raisers/intention/services/finance/accounts/utilities"
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

func (service *AccountsService) DeleteTransaction(
	ctx context.Context,
	request *contract_accounts.DeleteTransactionRequest,
) (*contract_accounts.DeleteTransactionResponse, error) {
	transaction_id := request.GetTransactionId()

	if transaction_id == 0 {
		return nil, fmt.Errorf("no transaction_id provided in deletion request")
	}

	if transaction_id < 0 {
		return nil, fmt.Errorf("transaction_id must be a positive value")
	}

	_ = resources.DeleteTransaction(service.db, uint(transaction_id))
	return &contract_accounts.DeleteTransactionResponse{}, status.Error(codes.NotFound, "")
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

func (service *AccountsService) UpdateTransaction(
	ctx context.Context,
	request *contract_accounts.UpdateTransactionRequest,
) (*contract_accounts.UpdateTransactionResponse, error) {
	transaction_pb := request.GetTransaction()

	if transaction_pb == nil {
		return nil, fmt.Errorf("no transaction provided")
	}

	if transaction_pb.GetTransactionId() <= 0 {
		return nil, fmt.Errorf("invalid transaction_id")
	}

	transaction := models.NewTransactionFromProto(transaction_pb)
	err := resources.UpdateTransaction(service.db, transaction)
	if err != nil {
		return nil, err
	}

	return &contract_accounts.UpdateTransactionResponse{}, nil
}
