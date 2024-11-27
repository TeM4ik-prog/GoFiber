package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("../server/database/dev.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	log.Println("Database connection established.")
}
