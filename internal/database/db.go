package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := "host=localhost user=host=localhost port=5432 user=postgres password=postgres dbname=golangwebserver sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	log.Println("Database connected")
}
