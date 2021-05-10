package db

import (
	// "os"
	"fmt"
	"log"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"

	"github.com/Oasixer/ShopifyChallenge2021/config"
)

// DB *grom.DB
type DB struct {
	*gorm.DB
}

// ConnectDB : connecting DB
func ConnectDB() (*DB, error) {
	dboptions := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.CONFIG.DB.Host,
		config.CONFIG.DB.Port,
		config.CONFIG.DB.User,
		config.CONFIG.DB.DBName,
		config.CONFIG.DB.Password,
		config.CONFIG.DB.SSLMode,
	)

	// connect to db
	db, err := gorm.Open(postgres.Open(dboptions),&gorm.Config{})

	if err != nil {
		log.Fatalf("Couldnt connect to db: %s", err)
	}

	return &DB{db}, nil
}
