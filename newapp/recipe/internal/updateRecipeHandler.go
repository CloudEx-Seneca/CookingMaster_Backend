package internal

import (
	"CookingMaster_Backend/newapp/common"
	"github.com/gin-gonic/gin"
)

// RecipeUpdateInput is the payload for updating a recipe.
type RecipeUpdateInput struct {
	RecipeID    uint     `json:"recipe_id" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Ingredients []string `json:"ingredients" binding:"required"`
	Image		string	 `json:"image" binding:"required"`
}

// UpdateRecipe updates the recipe and its ingredient associations.
// @Summary Update a recipe
// @Description Update recipe details and its ingredient associations.
// @Tags recipe
// @Accept json
// @Produce json
// @Param recipe body RecipeUpdateInput true "Recipe update payload"
// @Success 200 {object} common.Response
// @Security BearerAuth
// @Router /update [post]
func UpdateRecipe(c *gin.Context) {
	// This endpoint still requires JWT to ensure only the creator can update.
	userID := c.MustGet("user_id").(uint)
	var input RecipeUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		common.RespondJSON(c, 400, err.Error(), nil)
		return
	}

	var recipe Recipe
	if err := db.Where("id = ? AND user_id = ?", input.RecipeID, userID).Preload("Ingredients").First(&recipe).Error; err != nil {
		common.RespondJSON(c, 404, "Recipe not found", nil)
		return
	}

	// Update basic fields.
	recipe.Name = input.Name
	recipe.Description = input.Description

	// Process new ingredient list.
	newIngredients, err := findOrCreateIngredients(input.Ingredients)
	if err != nil {
		common.RespondJSON(c, 500, "Failed to process ingredients", nil)
		return
	}

	// Replace associations: delete old join records and set new ones.
	if err := db.Model(&recipe).Association("Ingredients").Replace(newIngredients); err != nil {
		common.RespondJSON(c, 500, "Failed to update ingredients", nil)
		return
	}

	if err := db.Save(&recipe).Error; err != nil {
		common.RespondJSON(c, 500, "Failed to update recipe", nil)
		return
	}
	common.RespondJSON(c, 200, "Recipe updated successfully", recipe)
}
