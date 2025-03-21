package internal

import (
	"CookingMaster_Backend/newapp/common"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

// initDB initializes the MySQL connection and performs auto-migration.
func InitDB() {
	db = common.ConnectMysql("recipe")
	// Auto-migrate Recipe and Ingredient models.
	if err := db.AutoMigrate(&Recipe{}, &Ingredient{}); err != nil {
		log.Fatal("AutoMigrate error: ", err)
	}
}
