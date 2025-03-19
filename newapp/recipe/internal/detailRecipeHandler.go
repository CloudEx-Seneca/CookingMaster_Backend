package internal

import (
	"CookingMaster_Backend/newapp/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RecipeDetail returns a single recipe (with ingredients) for the given recipe id.
// @Summary Get recipe detail
// @Description Get detailed information of a recipe by its ID.
// @Tags recipe
// @Produce json
// @Param id path int true "Recipe ID"  // Changed to path parameter
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 404 {object} common.Response
// @Router /detail/{id} [get]
func RecipeDetail(c *gin.Context) {
	// Get the 'id' from the URL path
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)  // Convert string ID to integer
	if err != nil {
		// If conversion fails, return a bad request error
		common.RespondJSON(c, http.StatusBadRequest, "Invalid Recipe ID format", nil)
		return
	}

	// Fetch the recipe from the database based on the ID
	var recipe Recipe
	if err := db.Where("id = ?", id).Preload("Ingredients").First(&recipe).Error; err != nil {
		// If no recipe is found, return a not found error
		common.RespondJSON(c, http.StatusNotFound, "Recipe not found", nil)
		return
	}

	// Return the recipe details
	common.RespondJSON(c, http.StatusOK, "Recipe detail retrieved successfully", recipe)
}
