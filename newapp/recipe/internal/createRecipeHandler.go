package internal

import "github.com/gin-gonic/gin"

// --- Handlers ---
// RecipeCreateInput is the expected payload for creating a recipe.
type RecipeCreateInput struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Ingredients []string `json:"ingredients" binding:"required"` // list of ingredient names
}

// CreateRecipe handles recipe creation.
// It uses JWT to set the UserID and creates the recipe along with its ingredient relations.
func CreateRecipe(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var input RecipeCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondJSON(c, 400, err.Error(), nil)
		return
	}

	// Get or create ingredient records.
	ingredients, err := findOrCreateIngredients(input.Ingredients)
	if err != nil {
		respondJSON(c, 500, "Failed to process ingredients", nil)
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
		respondJSON(c, 500, "Failed to create recipe", nil)
		return
	}
	respondJSON(c, 200, "Recipe created successfully", recipe)
}
