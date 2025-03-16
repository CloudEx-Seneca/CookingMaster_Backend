package internal

import (
	"CookingMaster_Backend/newapp/common"
	"github.com/gin-gonic/gin"
)

// ListRecipes returns all recipes in the database (publicly accessible).
func ListRecipes(c *gin.Context) {
	var recipes []Recipe
	if err := db.Preload("Ingredients").Find(&recipes).Error; err != nil {
		common.RespondJSON(c, 500, "Failed to fetch recipes", nil)
		return
	}
	common.RespondJSON(c, 200, "Recipe list", recipes)
}
