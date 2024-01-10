package db

import (
	"fmt"

	"gorm.io/gorm"
)

func Connect(dataSource *DataSource) *gorm.DB {
	fmt.Println("connecting to postgres database...")

	db, err := gorm.Open(dataSource.CreateDialector(), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v\n", err))
	}

	fmt.Println("successfully connected to database.")
	return db
}
