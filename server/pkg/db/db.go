package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	db_url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", 
	os.Getenv("DB_HOST"), 
	os.Getenv("DB_PORT"), 
	os.Getenv("DB_USER"), 
	os.Getenv("DB_NAME"), 
	os.Getenv("DB_PASSWORD"), 
	"disable")
	d, err := gorm.Open(postgres.Open(db_url), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}