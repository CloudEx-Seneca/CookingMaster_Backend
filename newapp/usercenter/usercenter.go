package main

import (
	"CookingMaster_Backend/newapp/common"
	"CookingMaster_Backend/newapp/usercenter/internal"
	"github.com/gin-gonic/gin"
	"os"
)

// ----- Main ----- //

func main() {
	internal.InitDB()
	router := gin.Default()
	router.Use(common.CorsMiddleware())

	// Group endpoints under /usercenter/v2.
	api := router.Group("/usercenter/v2")
	{
		// Registration and Login endpoints (POST only)
		api.POST("/register", internal.Register)
		api.POST("/login", internal.Login)

		// Profile endpoints (JWT protected)
		api.GET("/profile", common.JWTAuthMiddleware(), internal.GetProfile)
		api.POST("/profile", common.JWTAuthMiddleware(), internal.UpdateProfile)
	}

	addr := os.Getenv("SERVER_PORT")
	router.Run(addr)
}
