package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	DATABASE_URL = os.Getenv("DATABASE_URL")
)

func ConnectDatabase() {
	dsn := DATABASE_URL
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection failed!")
	}

	database.AutoMigrate(&Product{})

	DB = database
}
