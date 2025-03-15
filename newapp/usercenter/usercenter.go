package main

import (
	"CookingMaster_Backend/newapp/usercenter/internal"
	"github.com/gin-gonic/gin"
)

// ----- Main ----- //

func main() {
	internal.InitDB()
	router := gin.Default()

	// Group endpoints under /usercenter/v2.
	api := router.Group("/usercenter/v2")
	{
		// Registration and Login endpoints (POST only)
		api.POST("/register", internal.Register)
		api.POST("/login", internal.Login)

		// Profile endpoints (JWT protected)
		api.GET("/profile", internal.JWTAuthMiddleware(), internal.GetProfile)
		api.POST("/profile", internal.JWTAuthMiddleware(), internal.UpdateProfile)
	}

	// Run the server on port 8080.
	router.Run(":8080")
}
