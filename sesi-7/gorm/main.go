package main

import (
	"errors"
	"fmt"
	"gorm/database"
	"gorm/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()
	createUSer("leonardosiagian2001@gmail.com")
	getByUserID(1)
	updateUser(1, "leosiagiannn@gmail.com")
	createProduct(1, "FOO", "Laptop")
	getUsersWithProducts()
	deleteProductById(2)
}

func createUSer(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New User Data:", User)
}

func getByUserID(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User not found")
			return
		}
		println("Error getting user data:", err)
	}

	fmt.Println("User Data:", user)
}

func updateUser(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error
	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}
	fmt.Printf("Update user's email to %s\n", email)
}

func createProduct(userId uint, brand, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data:", err.Error())
		return
	}

	fmt.Println("New Product Data:", Product)
}

func getUsersWithProducts() {
	db := database.GetDB()

	users := models.User{}

	err := db.Preload("Product").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting users with products:", err)
		return
	}

	fmt.Println("Users Datas with Products")
	fmt.Printf("%+v\n", users)
}

func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}

	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product data:", err)
		return
	}

	fmt.Println("Deleted Product Data with id:", id)
}
