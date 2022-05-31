package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Create the databse struct
type Dbinstance struct {
	Db *gorm.DB
}

// Instantiate the database
var Database Dbinstance

// Create the connectDb function
func ConnectDb() {

	dsn := "host=192.168.95.41 user=jeffron password=root12 dbname=GoCrud port=5432 sslmode=disable TimeZone=Africa/Lagos"

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")

	db.Logger = logger.Default.LogMode(logger.Info)
}
