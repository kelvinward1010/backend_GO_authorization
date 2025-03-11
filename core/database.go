package core

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend_go/models"
	"backend_go/permissions"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Error when loading file .env: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Product{},
		&models.Role{},
		&models.Permission{},
	}

	err = db.AutoMigrate(modelsToMigrate...)
	if err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}

	permissions.SeedRolesAndPermissions(db)

	DB = db
	fmt.Println("✅ Database connected and migrated successfully!")
}
