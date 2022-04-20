package main

import (
	"Securing-Our-App-With-JSON-Web-Token/database"
	"Securing-Our-App-With-JSON-Web-Token/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":9000")
}
