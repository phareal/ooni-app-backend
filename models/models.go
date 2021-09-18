package models

import (
	"time"
)

type User struct {
	ID string `gorm:"primary_key"`
	Email string  `json:"email"`
	Password   string `json:"password"`
	CreatedAt time.Time
}

type Photos struct {
	ID string `gorm:"primary_key"`
	Url string
	IsUploaded bool
	User User  `gorm:"foreignKey:ID"`
	CreatedAt time.Time
}

type Admin struct {
	ID int `gorm:"primary_key, AUTO_INCREMENT"`
	Username string
	Password string
	CreatedAt time.Time
}
