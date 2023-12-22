package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	finance_models_pb "github.com/bar-raisers/intention/finance/models"
)

func TestNewTransaction(t *testing.T) {
	// Given
	expectedDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	expectedAccountId := "test_account_one"
	expectedPayeeId := "test_payee_one"
	expectedAmount := float32(123.45)

	// When
	transaction := NewTransaction(
		expectedDate,
		expectedAccountId,
		expectedPayeeId,
		expectedAmount,
	)

	// Then
	assert.Equal(t, expectedDate, transaction.Date)
	assert.Equal(t, expectedAccountId, transaction.AccountId)
	assert.Equal(t, expectedPayeeId, transaction.PayeeId)
	assert.Equal(t, expectedAmount, transaction.Amount)
}

func TestNewTransactionFromProto(t *testing.T) {
	// Given
	expectedDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	expectedAccountId := "test_account_one"
	expectedPayeeId := "test_payee_one"
	expectedAmount := float32(123.45)

	transactionProto := &finance_models_pb.Transaction{
		Date:      timestamppb.New(expectedDate),
		AccountId: expectedAccountId,
		PayeeId:   expectedPayeeId,
		Amount:    expectedAmount,
	}

	// When
	transaction := NewTransactionFromProto(transactionProto)

	// Then
	assert.Equal(t, expectedDate, transaction.Date)
	assert.Equal(t, expectedAccountId, transaction.AccountId)
	assert.Equal(t, expectedPayeeId, transaction.PayeeId)
	assert.Equal(t, expectedAmount, transaction.Amount)
}

func TestToProto(t *testing.T) {
	// Given
	expectedDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	expectedAccountId := "test_account_one"
	expectedPayeeId := "test_payee_one"
	expectedAmount := float32(123.45)

	transaction := NewTransaction(
		expectedDate,
		expectedAccountId,
		expectedPayeeId,
		expectedAmount,
	)

	// When
	transactionProto := transaction.ToProto()

	// Then
	assert.Equal(t, expectedDate, transactionProto.Date.AsTime())
	assert.Equal(t, expectedAccountId, transactionProto.AccountId)
	assert.Equal(t, expectedPayeeId, transactionProto.PayeeId)
	assert.Equal(t, expectedAmount, transactionProto.Amount)
}
