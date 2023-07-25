package config

import (
	"log"
	"os"
	"time"

	"chat/utils/errors"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	dbUserName := os.Getenv("DB_USERNAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbSSLmode := os.Getenv("DB_SSLMODE")

	metricsLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	dsn := "host=" + dbHost + " user=" + dbUserName + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSSLmode
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: metricsLogger})
	errors.ErrorCheck(err)
	log.Print("Database connected successfully!")
	db = d
}

func GetDB() *gorm.DB {
	return db
}
