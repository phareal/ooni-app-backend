package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitDatabase() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dsn := dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database")
	}
	SetupDatabase(db)
	return db
}
