package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JWT secret – in production, store securely.
var jwtSecret = []byte("mysecretkey")

// Response defines the standard API response structure.
type Response struct {
	Code    int         `json:"code"`    // e.g., 200 for success, 400/500 for errors
	Message string      `json:"message"` // descriptive message
	Data    interface{} `json:"data"`    // result payload (if any)
}

// respondJSON sends a JSON response with the given code, message, and data.
func respondJSON(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// --- Helper functions ---

// findOrCreateIngredients accepts a slice of ingredient names,
// fetches all ingredients in one query, and then compares in memory.
// Missing ingredients are bulk inserted.
func findOrCreateIngredients(ingredientNames []string) ([]Ingredient, error) {
	uniqueNames := make(map[string]bool)
	for _, name := range ingredientNames {
		n := strings.TrimSpace(strings.ToLower(name))
		if n != "" {
			uniqueNames[n] = true
		}
	}
	if len(uniqueNames) == 0 {
		return []Ingredient{}, nil
	}
	// Convert keys to slice.
	var names []string
	for name := range uniqueNames {
		names = append(names, name)
	}

	// Fetch all existing ingredients matching these names (case-insensitive).
	var existing []Ingredient
	if err := db.Where("LOWER(name) IN ?", names).Find(&existing).Error; err != nil {
		return nil, err
	}
	// Create a map for quick lookup.
	existingMap := make(map[string]Ingredient)
	for _, ing := range existing {
		existingMap[strings.ToLower(ing.Name)] = ing
	}

	// Identify missing names.
	var missing []Ingredient
	for name := range uniqueNames {
		if _, ok := existingMap[name]; !ok {
			missing = append(missing, Ingredient{Name: name})
		}
	}

	// Bulk insert missing ingredients.
	if len(missing) > 0 {
		if err := db.Create(&missing).Error; err != nil {
			return nil, err
		}
		// Add newly inserted ingredients to existingMap.
		for _, ing := range missing {
			existingMap[strings.ToLower(ing.Name)] = ing
		}
	}

	// Combine all ingredients from existingMap.
	var result []Ingredient
	for _, ing := range existingMap {
		result = append(result, ing)
	}
	return result, nil
}
