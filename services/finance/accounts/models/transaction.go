package models

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	finance_models_pb "github.com/bar-raisers/intention/models/finance"
)

type Transaction struct {
	gorm.Model

	Date      time.Time
	AccountId string
	PayeeId   string
	Amount    float32
}

func NewTransaction(
	transaction_id uint,
	date time.Time,
	accountId string,
	payeeId string,
	amount float32,
) *Transaction {
	return &Transaction{
		Model:     gorm.Model{ID: transaction_id},
		Date:      date,
		AccountId: accountId,
		PayeeId:   payeeId,
		Amount:    amount,
	}
}

func NewTransactionFromProto(transaction *finance_models_pb.Transaction) *Transaction {
	return NewTransaction(
		uint(transaction.GetTransactionId()),
		transaction.GetDate().AsTime(),
		transaction.GetAccountId(),
		transaction.GetPayeeId(),
		transaction.GetAmount(),
	)
}

func (t *Transaction) ToProto() *finance_models_pb.Transaction {
	return &finance_models_pb.Transaction{
		TransactionId: uint32(t.ID),
		Date:          timestamppb.New(t.Date),
		AccountId:     t.AccountId,
		PayeeId:       t.PayeeId,
		Amount:        t.Amount,
	}
}
