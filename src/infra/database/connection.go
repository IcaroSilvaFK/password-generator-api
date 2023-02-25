package infra

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DATABASE_USER     = "DATABASE_USER"
	DATABASE_PASSWORD = "DATABASE_PASSWORD"
	DATABASE_HOST     = "DATABASE_HOST"
	DATABASE_PORT     = "DATABASE_PORT"
	DATABASE_NAME     = "DATABASE_NAME"
)

func NewDatabase() *gorm.DB {
	var (
		user = os.Getenv(DATABASE_USER)
		pass = os.Getenv(DATABASE_PASSWORD)
		host = os.Getenv(DATABASE_HOST)
		port = os.Getenv(DATABASE_PORT)
		name = os.Getenv(DATABASE_NAME)
	)

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, name, port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connectionString,
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal(err)
	}

	return db

}
