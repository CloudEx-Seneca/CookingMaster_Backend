package internal

import "time"

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
