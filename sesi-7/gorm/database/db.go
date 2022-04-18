package database

import (
	"fmt"
	"gorm/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DB is the database connection
	db       *gorm.DB
	err      error
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "leonardo"
	dbname   = "learning-gorm"
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
