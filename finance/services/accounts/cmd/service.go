package main

import (
	"context"

	accounts "github.com/bar-raisers/intention/finance/contracts/accounts/v1"
)

type AccountsService struct{}

func NewAccountsService() *AccountsService {
	return &AccountsService{}
}

func (service *AccountsService) CreateTransaction(ctx context.Context, request *accounts.CreateTransactionRequest) (*accounts.CreateTransactionResponse, error) {
	return &accounts.CreateTransactionResponse{}, nil
}

func (service *AccountsService) ListTransactions(ctx context.Context, request *accounts.ListTransactionsRequest) (*accounts.ListTransactionsResponse, error) {
	return &accounts.ListTransactionsResponse{}, nil
}
