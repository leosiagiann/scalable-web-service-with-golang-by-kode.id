package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price string `json:"price"`
}

var CarDatas = []Car{}

func CreateCar(ctx *gin.Context) {
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas = append(CarDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var updateCar Car

	if err := ctx.ShouldBindJSON(&updateCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarDatas {
		if car.CarID == carID {
			condition = true
			CarDatas[i] = updateCar
			CarDatas[i].CarID = carID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Car not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Car with id %v has been updated", carID),
	})
}

func GetCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var getCar Car

	for i, c := range CarDatas {
		if c.CarID == carID {
			condition = true
			getCar = CarDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Car not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": getCar,
	})
}

func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carIndex int

	for i, c := range CarDatas {
		if c.CarID == carID {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Car not found",
		})
		return
	}

	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Car with id %v has been deleted", carID),
	})
}
