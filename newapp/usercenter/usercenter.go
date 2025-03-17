package main

import (
	"CookingMaster_Backend/newapp/common"
	docs "CookingMaster_Backend/newapp/usercenter/docs"
	"CookingMaster_Backend/newapp/usercenter/internal"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

// ----- Main ----- //
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	internal.InitDB()
	router := gin.Default()
	router.Use(common.CorsMiddleware())

	docs.SwaggerInfo.BasePath = "/usercenter/v2"

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	addr := os.Getenv("SERVER_PORT")
	router.Run(addr)
}
