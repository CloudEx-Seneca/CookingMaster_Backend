package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// JWT secret – in production, store securely.
var jwtSecret = []byte("mysecretkey")

var db *gorm.DB

// Response defines the standard API response structure.
type Response struct {
	Code    int         `json:"code"`    // e.g., 200 for success, 400/500 for errors
	Message string      `json:"message"` // descriptive message
	Data    interface{} `json:"data"`    // result payload (if any)
}

// Recipe represents a recipe entity.
type Recipe struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"size:255" json:"name"`
	Description string       `gorm:"size:1024" json:"description"`
	UserID      uint         `json:"user_id"` // creator of the recipe (from JWT when created)
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients;" json:"ingredients"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

// Ingredient represents an ingredient entity.
type Ingredient struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255;unique" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// RecipeSearchResponse is used in the search API to include missing ingredients.
type RecipeSearchResponse struct {
	Recipe
	MissingIngredients []string `json:"missing_ingredients"`
}

// initDB initializes the MySQL connection and performs auto-migration.
func initDB() {
	// Update DSN with your MySQL credentials.
	dsn := "root:haojiefu@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	// Auto-migrate Recipe and Ingredient models.
	if err := db.AutoMigrate(&Recipe{}, &Ingredient{}); err != nil {
		log.Fatal("AutoMigrate error: ", err)
	}
}

// respondJSON sends a JSON response with the given code, message, and data.
func respondJSON(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// JWTAuthMiddleware validates JWT tokens and sets "user_id" in context.
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			respondJSON(c, 401, "Authorization header required", nil)
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			respondJSON(c, 401, "Authorization header format must be Bearer {token}", nil)
			c.Abort()
			return
		}
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			respondJSON(c, 401, "Invalid token", nil)
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			respondJSON(c, 401, "Invalid token claims", nil)
			c.Abort()
			return
		}
		userID, ok := claims["user_id"].(float64)
		if !ok {
			respondJSON(c, 401, "Invalid token user_id", nil)
			c.Abort()
			return
		}
		c.Set("user_id", uint(userID))
		c.Next()
	}
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

// --- Request Payloads ---

// RecipeCreateInput is the expected payload for creating a recipe.
type RecipeCreateInput struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Ingredients []string `json:"ingredients" binding:"required"` // list of ingredient names
}

// RecipeUpdateInput is the payload for updating a recipe.
type RecipeUpdateInput struct {
	RecipeID    uint     `json:"recipe_id" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Ingredients []string `json:"ingredients" binding:"required"`
}

// RecipeDeleteInput is the payload for deleting a recipe.
type RecipeDeleteInput struct {
	RecipeID uint `json:"recipe_id" binding:"required"`
}

// RecipeDetailInput is the payload for recipe detail query.
type RecipeDetailInput struct {
	RecipeID uint `form:"recipe_id" binding:"required"` // using query parameter for GET
}

// RecipeSearchInput is for searching recipes.
type RecipeSearchInput struct {
	Ingredients string `form:"ingredients" binding:"required"` // comma-separated ingredient names
}

// --- Handlers ---

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

// ListRecipes returns all recipes in the database (publicly accessible).
func ListRecipes(c *gin.Context) {
	var recipes []Recipe
	if err := db.Preload("Ingredients").Find(&recipes).Error; err != nil {
		respondJSON(c, 500, "Failed to fetch recipes", nil)
		return
	}
	respondJSON(c, 200, "Recipe list", recipes)
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

// UpdateRecipe updates the recipe and its ingredient associations.
func UpdateRecipe(c *gin.Context) {
	// This endpoint still requires JWT to ensure only the creator can update.
	userID := c.MustGet("user_id").(uint)
	var input RecipeUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondJSON(c, 400, err.Error(), nil)
		return
	}

	var recipe Recipe
	if err := db.Where("id = ? AND user_id = ?", input.RecipeID, userID).Preload("Ingredients").First(&recipe).Error; err != nil {
		respondJSON(c, 404, "Recipe not found", nil)
		return
	}

	// Update basic fields.
	recipe.Name = input.Name
	recipe.Description = input.Description

	// Process new ingredient list.
	newIngredients, err := findOrCreateIngredients(input.Ingredients)
	if err != nil {
		respondJSON(c, 500, "Failed to process ingredients", nil)
		return
	}

	// Replace associations: delete old join records and set new ones.
	if err := db.Model(&recipe).Association("Ingredients").Replace(newIngredients); err != nil {
		respondJSON(c, 500, "Failed to update ingredients", nil)
		return
	}

	if err := db.Save(&recipe).Error; err != nil {
		respondJSON(c, 500, "Failed to update recipe", nil)
		return
	}
	respondJSON(c, 200, "Recipe updated successfully", recipe)
}

// DeleteRecipe deletes a recipe and clears its join table connections.
func DeleteRecipe(c *gin.Context) {
	// This endpoint requires JWT to ensure only the creator can delete.
	userID := c.MustGet("user_id").(uint)
	var input RecipeDeleteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondJSON(c, 400, err.Error(), nil)
		return
	}

	var recipe Recipe
	if err := db.Where("id = ? AND user_id = ?", input.RecipeID, userID).Preload("Ingredients").First(&recipe).Error; err != nil {
		respondJSON(c, 404, "Recipe not found", nil)
		return
	}

	// Clear associations in the join table.
	if err := db.Model(&recipe).Association("Ingredients").Clear(); err != nil {
		respondJSON(c, 500, "Failed to clear recipe ingredients", nil)
		return
	}
	// Delete the recipe.
	if err := db.Delete(&recipe).Error; err != nil {
		respondJSON(c, 500, "Failed to delete recipe", nil)
		return
	}
	respondJSON(c, 200, "Recipe deleted successfully", nil)
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

// --- Main ---

func main() {
	initDB()
	router := gin.Default()

	// Recipe API endpoints group with prefix /recipe/v2.
	api := router.Group("/recipe/v2")
	{
		// Endpoints that require authentication.
		api.POST("/create", JWTAuthMiddleware(), CreateRecipe)
		api.POST("/update", JWTAuthMiddleware(), UpdateRecipe)
		api.POST("/delete", JWTAuthMiddleware(), DeleteRecipe)

		// Public endpoints for reading recipes.
		api.GET("/list", ListRecipes)
		api.GET("/detail", RecipeDetail)

		// Public search endpoint.
		api.GET("/search", SearchRecipe)
	}

	// Run the server on port 8080.
	router.Run(":8080")
}
