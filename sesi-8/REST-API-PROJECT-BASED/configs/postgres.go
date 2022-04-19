package configs

import (
	"fmt"
	"sesi8-project/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	host   string = "localhost"
	port   int    = 5432
	user   string = "postgres"
	pass   string = "leonardo"
	dbname string = "sesi-8"
)

func NewPostgres() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if !db.HasTable(models.Person{}) {
		db.AutoMigrate(models.Person{})
	}

	return db
}
