package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectDB(dbName string) *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	mysqlPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	dsn := fmt.Sprintf("root:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	return db
}
