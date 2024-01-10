package resources

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/bar-raisers/intention/finance/services/accounts/models"
)

func Migrate(db *gorm.DB) {
	fmt.Println("migrating database tables...")

	err := db.AutoMigrate(&models.Transaction{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate database tables: %v\n", err))
	}

	fmt.Println("successfully migrated database tables.")
}
