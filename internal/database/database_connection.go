package database

import (
	"fmt"
	"log"
	"user-profile-service/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnection() {
	dbHost := config.GetValue("DATABASE_HOST")
	dbPort := config.GetValue("DATABASE_PORT")
	dbUser := config.GetValue("DATABASE_USER")
	dbPassword := config.GetValue("DATABASE_PASSWORD")
	dbName := config.GetValue("DATABASE_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	DB = db
	log.Print("Connect to database success")
}
