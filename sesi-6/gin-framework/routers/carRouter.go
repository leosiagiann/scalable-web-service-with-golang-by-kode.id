package routers

import (
	"git-framework/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	r.GET("/", controllers.GetCars)

	return r
}
