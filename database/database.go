package database

import (
	"fmt"
	"log"
	"os"

	"github.com/greybluesea/jwt_mvc_on_aws/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	*gorm.DB
}

var DB DBInstance

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=verify-full sslrootcert=%s TimeZone=Pacific/Auckland", os.Getenv("DB_ENDPOINT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("SSL_ROOT_CERT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
		os.Exit(2)
	}

	log.Println("DB connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.User{})

	DB = DBInstance{
		db,
	}

}
