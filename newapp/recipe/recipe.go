package main

import (
	"CookingMaster_Backend/newapp/common"
	"CookingMaster_Backend/newapp/recipe/internal"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	internal.InitDB()
	router := gin.Default()
	router.Use(common.CorsMiddleware())

	// Recipe API endpoints group with prefix /recipe/v2.
	api := router.Group("/recipe/v2")
	{
		// Endpoints that require authentication.
		api.POST("/create", common.JWTAuthMiddleware(), internal.CreateRecipe)
		api.POST("/update", common.JWTAuthMiddleware(), internal.UpdateRecipe)
		api.POST("/delete", common.JWTAuthMiddleware(), internal.DeleteRecipe)

		// Public endpoints for reading recipes.
		api.GET("/list", internal.ListRecipes)
		api.GET("/detail", internal.RecipeDetail)

		// Public search endpoint.
		api.GET("/search", internal.SearchRecipe)
	}

	addr := os.Getenv("SERVER_PORT")
	router.Run(addr)
}
