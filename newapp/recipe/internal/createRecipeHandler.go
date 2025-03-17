package internal

import (
	"CookingMaster_Backend/newapp/common"
	"github.com/gin-gonic/gin"
)

// --- Handlers ---
// RecipeCreateInput is the expected payload for creating a recipe.
type RecipeCreateInput struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Ingredients []string `json:"ingredients" binding:"required"` // list of ingredient names
}

// CreateRecipe handles recipe creation.
// It uses JWT to set the UserID and creates the recipe along with its ingredient relations.
// @Summary Create a new recipe
// @Description Create a recipe with name, description and list of ingredients. UserID is set from JWT.
// @Tags recipe
// @Accept json
// @Produce json
// @Param recipe body RecipeCreateInput true "Recipe creation payload"
// @Success 200 {object} common.Response
// @Security BearerAuth
// @Router /create [post]
func CreateRecipe(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var input RecipeCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		common.RespondJSON(c, 400, err.Error(), nil)
		return
	}

	// Get or create ingredient records.
	ingredients, err := findOrCreateIngredients(input.Ingredients)
	if err != nil {
		common.RespondJSON(c, 500, "Failed to process ingredients", nil)
		return
	}

	// Create the recipe with associations.
	recipe := Recipe{
		Name:        input.Name,
		Description: input.Description,
		UserID:      userID,
		Ingredients: ingredients,
	}
	if err := db.Create(&recipe).Error; err != nil {
		common.RespondJSON(c, 500, "Failed to create recipe", nil)
		return
	}
	common.RespondJSON(c, 200, "Recipe created successfully", recipe)
}
