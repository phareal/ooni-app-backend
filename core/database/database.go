package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitDatabase() *gorm.DB {
	dbUser := os.Getenv("QOVERY_MYSQL_ZD1EF4E14_LOGIN")
	dbPass := os.Getenv("QOVERY_MYSQL_ZD1EF4E14_PASS")
	dbName := os.Getenv("QOVERY_MYSQL_ZD1EF4E14_DEFAULT_DATABASE_NAME")
	dbHost := os.Getenv("QOVERY_MYSQL_ZD1EF4E14_HOST")
	dsn := dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database")
	}
	SetupDatabase(db)
	return db
}


func CloseConnection(db *gorm.DB) {

}
