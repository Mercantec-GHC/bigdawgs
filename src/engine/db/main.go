package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsu := os.Getenv("DATABASE_URL")
	fmt.Printf("DATABASE_URL: %s\n", dsu)
	if dsu == "" {
		log.Println("DATABASE_URL is not set")
		return nil, os.ErrNotExist
	}

	db, err := gorm.Open(postgres.Open(dsu), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	bdDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := bdDB.Ping(); err != nil {
		return nil, err
	}

	log.Println("database connection established")
	return db, nil
}
