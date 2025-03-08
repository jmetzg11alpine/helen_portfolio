package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"log"

	"helen-portfolio/backend/models"
)

var DB *gorm.DB

func Connect() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("data/helen-portfolio.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	DB.AutoMigrate(&models.BlogPost{}, &models.BlogComment{}, &models.Merch{}, &models.MerchHistory{})

	seedDB(DB)
	return nil
}
