package models

import "time"

type User struct {
	ID        uint   `gorm:"primary_key"`
	Email     string `gorm:"not null;unique;type:varchar(191)"`
	Product   []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}