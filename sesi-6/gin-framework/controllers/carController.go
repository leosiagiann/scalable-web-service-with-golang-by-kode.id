package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	MOdel string `json:"model"`
	PRice int    `json:"price"`
}

var cars []Car

func createCar(c *gin.Context) {
	var car Car

	if err := c.ShouldBindJSON(&car); err != nil {
		c.AbortWithError(400, err)
		return
	}

	car.CarID = "car-" + car.Brand + "-" + car.MOdel

}

func GetCars(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"cars": cars,
	})
}
