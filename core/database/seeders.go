package database

import (
	"fmt"
	"gorm.io/gorm"
	"ooni/backend/models"
)

func SeedDatabase(db *gorm.DB)  {
	 seedAdminTable(db)
	 seedUserTable(db)
}

func seedAdminTable(db *gorm.DB)  {
	db.Create(&models.Admin{
		Username: "admin",
		Password: "qP2gMH-Lb$CKcGn",
	})
}

func seedUserTable(db *gorm.DB) {
	for i := 0; i < 5; i++ {
		db.Create(&models.User{
			 Email: fmt.Sprint("john",i)+"@doe.com",
			 Password: "qP2gMH",
		})
	}
}