package models

import (
	"time"
)

type User struct {
	ID int `gorm:"primary_key, AUTO_INCREMENT"`
	Email string
	Password   string
	CreatedAt time.Time
}

type Photos struct {
	ID int `gorm:"primary_key, AUTO_INCREMENT"`
	Url string
	IsUploaded bool
	UserID int
	User User
	CreatedAt time.Time
}

type Admin struct {
	ID int `gorm:"primary_key, AUTO_INCREMENT"`
	Username string
	Password string
	CreatedAt time.Time
}
