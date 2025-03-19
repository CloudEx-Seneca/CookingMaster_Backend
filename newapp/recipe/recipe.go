package main

import (
	"CookingMaster_Backend/newapp/common"
	docs "CookingMaster_Backend/newapp/recipe/docs"
	"CookingMaster_Backend/newapp/recipe/internal"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	internal.InitDB()
	router := gin.Default()
	router.Use(common.CorsMiddleware())

	docs.SwaggerInfo.BasePath = "/recipe/v2"

	// Recipe API endpoints group with prefix /recipe/v2.
	api := router.Group("/recipe/v2")
	{
		// Endpoints that require authentication.
		api.POST("/create", common.JWTAuthMiddleware(), internal.CreateRecipe)
		api.POST("/update", common.JWTAuthMiddleware(), internal.UpdateRecipe)
		api.POST("/delete", common.JWTAuthMiddleware(), internal.DeleteRecipe)

		// Public endpoints for reading recipes.
		api.GET("/list", internal.ListRecipes)
		api.GET("/detail/:id", internal.RecipeDetail) 

		// Public search endpoint.
		api.GET("/search", internal.SearchRecipe)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	addr := os.Getenv("SERVER_PORT")
	router.Run(addr)
}
