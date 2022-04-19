package main

import (
	"assigment-2/config"
	"assigment-2/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	_port_ = ":9000"
)

func main() {
	db := config.StartDB()
	if db != nil {
		fmt.Println("Running database is failed")
	}

	orderController := controllers.NewControllerOrder(db)
	itemController := controllers.NewControllerItem(db)

	router := gin.Default()

	router.GET("/orders", orderController.GetOrder)
	router.POST("/orders", orderController.CreateOrder)
	router.PUT("/orders/:orderId", orderController.UpdateOrder)
	router.DELETE("/orders/:orderId", orderController.DeleteOrder)

	router.GET("/items", itemController.GetItem)
	router.POST("/items", itemController.CreateItem)
	router.PUT("/items/:itemId", itemController.UpdateItem)
	router.DELETE("/items/:itemId", itemController.DeleteItem)
	router.Run(_port_)
}
