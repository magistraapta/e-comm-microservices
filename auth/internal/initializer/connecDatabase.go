package initializer

import (
	"auth/internal/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {

	databaseConfig := os.Getenv("DATABASE_CONFIG")
	db, err := gorm.Open(postgres.Open(databaseConfig), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.User{})

	return db
}
