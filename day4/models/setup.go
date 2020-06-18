// models/setup.go

package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("postgres", "host=summerclub20-postgres.vinnoti.com port=5432 user=summerclub2020 dbname=christian_mahardhika password=Multatuli1820 sslmode=disable")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Car{})

	DB = database

}
