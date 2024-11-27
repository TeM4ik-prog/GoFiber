package database

import (
	"fiber-app/database/models"
	"log"
)

// AutoMigrate выполняет миграции
func AutoMigrate() {
	if DB == nil {
		log.Fatal("Database is not initialized. Call Connect() before AutoMigrate().")
	}


	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed.")
}
