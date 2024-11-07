package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func Connect() {
	var err error
	dsn := "host=localhost user=postgres password=dokuzu_desu dbname=test port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}

func Disconnect() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatal("Failed to close the database connection:", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
