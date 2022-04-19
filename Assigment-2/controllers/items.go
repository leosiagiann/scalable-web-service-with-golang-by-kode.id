package controllers

import (
	"assigment-2/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDBI struct {
	DB *gorm.DB
}

func NewControllerItem(db *gorm.DB) *InDBI {
	return &InDBI{
		DB: db,
	}
}

func (in *InDBI) GetItem(c *gin.Context) {
	var (
		item   []models.Item
		result gin.H
	)

	err := in.DB.Find(&item).Error
	if err != nil {
		result = gin.H{
			"result": nil,
			"error":  err.Error(),
		}
	}

	if len(item) <= 0 {
		result = gin.H{
			"result": nil,
			"error":  "Data is empty",
		}
	} else {
		result = gin.H{
			"data": item,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (in *InDBI) CreateItem(c *gin.Context) {
	var item models.Item

	err := json.NewDecoder(c.Request.Body).Decode(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.DB.Create(&item).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": item,
	})
}

func (in *InDBI) UpdateItem(c *gin.Context) {
	var (
		item    models.Item
		newItem models.Item
	)
	id := c.Param("itemId")

	err := in.DB.First(&item, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&newItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.DB.Model(&item).Updates(newItem).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newItem,
	})
}

func (in *InDBI) DeleteItem(c *gin.Context) {
	var (
		item models.Item
	)
	id := c.Param("itemId")

	err := in.DB.First(&item, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.DB.Delete(&item).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "delete data success !",
	})
}
