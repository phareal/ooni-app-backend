package database

import (
	"gorm.io/gorm"
	"ooni/backend/models"
)

func SeedDatabase(db *gorm.DB)  {
	 seedAdminTable(db)
}

func seedAdminTable(db *gorm.DB)  {
	db.Create(&models.Admin{
		Username: "admin",
		Password: "qP2gMH-Lb$CKcGn",
	})
}