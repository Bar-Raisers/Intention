package resources

import (
	"gorm.io/gorm"

	"github.com/bar-raisers/intention/services/finance/accounts/models"
)

func CreateTransaction(db *gorm.DB, transaction *models.Transaction) error {
	result := db.Create(transaction)
	if result != nil {
		return result.Error
	}
	return nil
}

func DeleteTransaction(db *gorm.DB, transaction_id uint) error {
	return db.Delete(&models.Transaction{}, transaction_id).Error
}

func ListTransactions(db *gorm.DB) []*models.Transaction {
	transactions := []*models.Transaction{}
	db.Find(&transactions)
	return transactions
}

func UpdateTransaction(db *gorm.DB, transaction *models.Transaction) error {
	return db.Model(&transaction).Updates(&transaction).Error
}
