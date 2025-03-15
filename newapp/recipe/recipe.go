package main

import (
	"CookingMaster_Backend/newapp/recipe/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	internal.InitDB()
	router := gin.Default()

	// Recipe API endpoints group with prefix /recipe/v2.
	api := router.Group("/recipe/v2")
	{
		// Endpoints that require authentication.
		api.POST("/create", internal.JWTAuthMiddleware(), internal.CreateRecipe)
		api.POST("/update", internal.JWTAuthMiddleware(), internal.UpdateRecipe)
		api.POST("/delete", internal.JWTAuthMiddleware(), internal.DeleteRecipe)

		// Public endpoints for reading recipes.
		api.GET("/list", internal.ListRecipes)
		api.GET("/detail", internal.RecipeDetail)

		// Public search endpoint.
		api.GET("/search", internal.SearchRecipe)
	}

	// Run the server on port 8080.
	router.Run(":8080")
}
