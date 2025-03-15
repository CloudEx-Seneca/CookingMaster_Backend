package internal

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// RecipeSearchInput is for searching recipes.
type RecipeSearchInput struct {
	Ingredients string `form:"ingredients" binding:"required"` // comma-separated ingredient names
}

// RecipeSearchResponse is used in the search API to include missing ingredients.
type RecipeSearchResponse struct {
	Recipe
	MissingIngredients []string `json:"missing_ingredients"`
}

// SearchRecipe searches recipes by a list of ingredient names provided as a comma-separated string.
// It returns recipes that have at least one matching ingredient along with a list of missing ingredients.
func SearchRecipe(c *gin.Context) {
	var input RecipeSearchInput
	if err := c.ShouldBindQuery(&input); err != nil {
		respondJSON(c, 400, err.Error(), nil)
		return
	}
	// Split and normalize the input ingredient names.
	inputIngredients := strings.Split(input.Ingredients, ",")
	availableMap := make(map[string]bool)
	for _, ing := range inputIngredients {
		name := strings.TrimSpace(strings.ToLower(ing))
		if name != "" {
			availableMap[name] = true
		}
	}

	var recipes []Recipe
	if err := db.Preload("Ingredients").Find(&recipes).Error; err != nil {
		respondJSON(c, 500, "Failed to fetch recipes", nil)
		return
	}

	var result []RecipeSearchResponse
	for _, recipe := range recipes {
		matched := false
		var missing []string
		for _, ing := range recipe.Ingredients {
			ingName := strings.ToLower(ing.Name)
			if availableMap[ingName] {
				matched = true
			} else {
				missing = append(missing, ing.Name)
			}
		}
		if matched {
			result = append(result, RecipeSearchResponse{
				Recipe:             recipe,
				MissingIngredients: missing,
			})
		}
	}
	respondJSON(c, 200, "Search completed", result)
}
