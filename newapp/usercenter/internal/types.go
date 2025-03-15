package internal

import "time"

// Response defines the professional JSON response format.
type Response struct {
	Code    int         `json:"code"`    // e.g., 200 for success, 400/500 for errors
	Message string      `json:"message"` // descriptive message
	Data    interface{} `json:"data"`    // result payload (if any)
}

// User entity definition.
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"size:255;unique" json:"email"`
	Password  string    `gorm:"size:255" json:"-"` // never expose password
	Nickname  string    `gorm:"size:255" json:"nickname"`
	Sex       string    `gorm:"size:10" json:"sex"`
	Info      string    `gorm:"size:1024" json:"info"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
