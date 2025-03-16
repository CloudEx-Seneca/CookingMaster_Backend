package internal

import (
	"CookingMaster_Backend/newapp/common"
	"gorm.io/gorm"
)

// Global DB variable
var db *gorm.DB

// initDB connects to MySQL and performs auto-migration.
func InitDB() {
	db = common.ConnectMysql("usercenter")
	// Auto-migrate the User table.
	db.AutoMigrate(&User{})
}
