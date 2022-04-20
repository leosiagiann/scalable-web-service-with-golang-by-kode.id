package database

import (
	"Securing-Our-App-With-JSON-Web-Token/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DB is the database connection
	db     *gorm.DB
	err    error
	host   = "localhost"
	user   = "postgres"
	pass   = "leonardo"
	dbport = "5432"
	dbname = "simple-api"
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, pass, dbport, dbname)
	dsn := config

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
	}

	db.AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
