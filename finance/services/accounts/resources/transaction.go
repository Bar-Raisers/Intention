package resources

import (
	"gorm.io/gorm"

	"github.com/bar-raisers/intention/finance/services/accounts/models"
)

func CreateTransaction(db *gorm.DB, transaction *models.Transaction) error {
	result := db.Create(transaction)
	if result != nil {
		return result.Error
	}

	return nil
}

func ListTransactions(db *gorm.DB) []*models.Transaction {
	transactions := []*models.Transaction{}
	db.Find(&transactions)
	return transactions
}
