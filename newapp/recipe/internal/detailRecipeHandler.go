package internal

import "github.com/gin-gonic/gin"

// RecipeDetailInput is the payload for recipe detail query.
type RecipeDetailInput struct {
	RecipeID uint `form:"recipe_id" binding:"required"` // using query parameter for GET
}

// RecipeDetail returns a single recipe (with ingredients) for the given recipe id.
func RecipeDetail(c *gin.Context) {
	var input RecipeDetailInput
	if err := c.ShouldBindQuery(&input); err != nil {
		respondJSON(c, 400, err.Error(), nil)
		return
	}

	var recipe Recipe
	if err := db.Where("id = ?", input.RecipeID).Preload("Ingredients").First(&recipe).Error; err != nil {
		respondJSON(c, 404, "Recipe not found", nil)
		return
	}
	respondJSON(c, 200, "Recipe detail", recipe)
}
