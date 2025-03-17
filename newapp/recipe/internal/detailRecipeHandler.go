package internal

import (
	"CookingMaster_Backend/newapp/common"
	"github.com/gin-gonic/gin"
)

// RecipeDetailInput is the payload for recipe detail query.
type RecipeDetailInput struct {
	RecipeID uint `form:"recipe_id" binding:"required"` // using query parameter for GET
}

// RecipeDetail returns a single recipe (with ingredients) for the given recipe id.
// @Summary Get recipe detail
// @Description Get detailed information of a recipe by its ID.
// @Tags recipe
// @Produce json
// @Param recipe_id query int true "Recipe ID"
// @Success 200 {object} common.Response
// @Router /detail [get]
func RecipeDetail(c *gin.Context) {
	var input RecipeDetailInput
	if err := c.ShouldBindQuery(&input); err != nil {
		common.RespondJSON(c, 400, err.Error(), nil)
		return
	}

	var recipe Recipe
	if err := db.Where("id = ?", input.RecipeID).Preload("Ingredients").First(&recipe).Error; err != nil {
		common.RespondJSON(c, 404, "Recipe not found", nil)
		return
	}
	common.RespondJSON(c, 200, "Recipe detail", recipe)
}
