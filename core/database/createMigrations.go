package database

import (
	"gorm.io/gorm"
	"ooni/backend/models"
)

func SetupDatabase (db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Photos{}, &models.Admin{})
	if err != nil {
		panic("We cannot migrate your database")
		return
	}
}
