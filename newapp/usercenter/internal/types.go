package internal

import "time"

// User entity definition.
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"size:255;unique" json:"email"`
	Password  string    `gorm:"size:255" json:"-"` // never expose password
	Nickname  string    `gorm:"size:255" json:"nickname"`
	Sex       string    `gorm:"size:10" json:"sex"`
	Info      string    `gorm:"size:1024" json:"info"`
	AvatarUrl string	`gorm:"size:2048" json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
