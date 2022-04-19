package main

import (
	"latihan/controllers"
	"latihan/helper"

	"github.com/gin-gonic/gin"
)

func main() {
	PORT := ":9000"
	p := helper.Post{}
	controllerPost := controllers.InPost{
		Post: &p,
	}

	router := gin.Default()

	router.GET("/posts", controllerPost.GetAllPosts)
	router.GET("/posts/:id", controllerPost.GetPostById)
	router.POST("/posts", controllerPost.CreateNewPost)

	router.Run(PORT)
}
