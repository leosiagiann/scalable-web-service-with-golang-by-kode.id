package main

import (
	"assigment-2/config"
	"assigment-2/controllers"
	"fmt"

	_ "assigment-2/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const (
	_port_ = ":9000"
)

// @title Assigment-2
// @version 1.0
// @description This is a sample server Petstore server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9000
// @BasePath /
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

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(_port_)
}
