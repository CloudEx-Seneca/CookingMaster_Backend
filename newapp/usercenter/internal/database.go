package internal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// Global DB variable
var db *gorm.DB

// initDB connects to MySQL and performs auto-migration.
func InitDB() {
	// Update DSN with your MySQL credentials and database details.
	dsn := "root:haojiefu@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	// Auto-migrate the User table.
	db.AutoMigrate(&User{})
}
