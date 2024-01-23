package utilities

import (
	"github.com/bar-raisers/intention/finance/services/accounts/models"
	pb_finance_models "github.com/bar-raisers/intention/models/finance"
)

func ConvertTransactionsToProto(transactions []*models.Transaction) []*pb_finance_models.Transaction {
	transactions_pb := make([]*pb_finance_models.Transaction, 0)
	for _, transaction := range transactions {
		transactions_pb = append(transactions_pb, transaction.ToProto())
	}

	return transactions_pb
}
