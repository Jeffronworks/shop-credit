package database

import (
	"log"
	"os"

	"github.com/jeffronworks/shop-credit/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var Database Dbinstance

func ConnectDb() {
	dsn := "host=192.168.95.41 user=jeffron password=root12 dbname=shopcredit port=5432 sslmode=disable TimeZone=Africa/Lagos"

	// connect to postges DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect to DB \n", err.Error())
		os.Exit(2)
	}

	log.Println("Successfully connected to database")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Runing Migration")

	// Add migrations

	db.AutoMigrate(&models.Order{}, &models.Product{}, &models.User{})

	Database = Dbinstance{Db: db}
}
