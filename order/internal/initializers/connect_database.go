package initializers

import (
	"log"
	"order/internal/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {

	datbaaseConfig := os.Getenv("DATABASE_CONFIG")
	db, err := gorm.Open(postgres.Open(datbaaseConfig), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&model.Order{})

	return db
}
