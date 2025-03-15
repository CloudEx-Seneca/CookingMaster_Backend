package internal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

// initDB initializes the MySQL connection and performs auto-migration.
func InitDB() {
	// Update DSN with your MySQL credentials.
	dsn := "root:haojiefu@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	// Auto-migrate Recipe and Ingredient models.
	if err := db.AutoMigrate(&Recipe{}, &Ingredient{}); err != nil {
		log.Fatal("AutoMigrate error: ", err)
	}
}
