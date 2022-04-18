package main

import (
	"database/sql"
	"fmt"
	"go-sql/database"
	"go-sql/functions"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func main() {
	db := database.CreateConnection()

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
	functions.CreateEmployee(db)
	functions.DeleteEmplyoee(db)
	functions.UpdateEmployee(db)
	functions.GetEmployees(db)
}
