package internal

import (
	"CookingMaster_Backend/newapp/common"
	"github.com/gin-gonic/gin"
)

// RecipeDeleteInput is the payload for deleting a recipe.
type RecipeDeleteInput struct {
	RecipeID uint `json:"recipe_id" binding:"required"`
}

// DeleteRecipe deletes a recipe and clears its join table connections.
func DeleteRecipe(c *gin.Context) {
	// This endpoint requires JWT to ensure only the creator can delete.
	userID := c.MustGet("user_id").(uint)
	var input RecipeDeleteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		common.RespondJSON(c, 400, err.Error(), nil)
		return
	}

	var recipe Recipe
	if err := db.Where("id = ? AND user_id = ?", input.RecipeID, userID).Preload("Ingredients").First(&recipe).Error; err != nil {
		common.RespondJSON(c, 404, "Recipe not found", nil)
		return
	}

	// Clear associations in the join table.
	if err := db.Model(&recipe).Association("Ingredients").Clear(); err != nil {
		common.RespondJSON(c, 500, "Failed to clear recipe ingredients", nil)
		return
	}
	// Delete the recipe.
	if err := db.Delete(&recipe).Error; err != nil {
		common.RespondJSON(c, 500, "Failed to delete recipe", nil)
		return
	}
	common.RespondJSON(c, 200, "Recipe deleted successfully", nil)
}
