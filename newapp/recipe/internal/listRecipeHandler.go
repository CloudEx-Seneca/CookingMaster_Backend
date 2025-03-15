package internal

import "github.com/gin-gonic/gin"

// ListRecipes returns all recipes in the database (publicly accessible).
func ListRecipes(c *gin.Context) {
	var recipes []Recipe
	if err := db.Preload("Ingredients").Find(&recipes).Error; err != nil {
		respondJSON(c, 500, "Failed to fetch recipes", nil)
		return
	}
	respondJSON(c, 200, "Recipe list", recipes)
}
