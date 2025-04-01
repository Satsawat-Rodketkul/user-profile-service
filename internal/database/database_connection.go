package database

import (
	"context"
	"fmt"
	"log"
	"time"
	"user-profile-service/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n===============================\n", sql)
}

var DB *gorm.DB

func DBconnection() {
	dbHost := config.GetValue("DATABASE_HOST")
	dbPort := config.GetValue("DATABASE_PORT")
	dbUser := config.GetValue("DATABASE_USER")
	dbPassword := config.GetValue("DATABASE_PASSWORD")
	dbName := config.GetValue("DATABASE_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	DB = db
	log.Print("Connect to database success")
}
