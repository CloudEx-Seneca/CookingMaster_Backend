package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// JWT secret key – in production, store this securely
var jwtSecret = []byte("mysecretkey")

// Global DB variable
var db *gorm.DB

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

// initDB connects to MySQL and performs auto-migration.
func initDB() {
	// Update DSN with your MySQL credentials and database details.
	dsn := "root:haojiefu@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	// Auto-migrate the User table.
	db.AutoMigrate(&User{})
}

// respondJSON sends a JSON response with the specified code, message, and data.
func respondJSON(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// hashPassword hashes the plaintext password using bcrypt.
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPasswordHash compares a plaintext password with its hashed version.
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// generateJWT generates a JWT token with user_id and email claims.
func generateJWT(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(72 * time.Hour).Unix(), // token expires in 72 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWTAuthMiddleware validates the token and sets the user_id in the context.
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
			// Validate signing method.
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

// ----- Handlers ----- //

// RegisterInput defines the expected parameters for user registration.
type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Register handles user registration.
// If the email has been registered before, it returns an error message.
func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondJSON(c, 400, err.Error(), nil)
		return
	}

	// Check if the email is already registered.
	var user User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err == nil {
		respondJSON(c, 400, "Email already registered", nil)
		return
	}

	// Hash password and create user.
	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		respondJSON(c, 500, "Failed to hash password", nil)
		return
	}
	user = User{
		Email:    input.Email,
		Password: hashedPassword,
	}
	if err := db.Create(&user).Error; err != nil {
		respondJSON(c, 500, "Failed to create user", nil)
		return
	}
	respondJSON(c, 200, "Register successful", nil)
}

// LoginInput defines the expected parameters for user login.
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Login handles user login.
// It verifies whether the email is registered, checks the password, and returns a JWT on success.
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondJSON(c, 400, err.Error(), nil)
		return
	}

	var user User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		respondJSON(c, 400, "Email not registered", nil)
		return
	}

	if !checkPasswordHash(input.Password, user.Password) {
		respondJSON(c, 400, "Password wrong", nil)
		return
	}

	token, err := generateJWT(user.ID, user.Email)
	if err != nil {
		respondJSON(c, 500, "Failed to generate token", nil)
		return
	}
	respondJSON(c, 200, "Login successful", gin.H{"token": token})
}

// GetProfile returns the profile details of the authenticated user.
func GetProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		respondJSON(c, 404, "User not found", nil)
		return
	}
	// Exclude the password field.
	user.Password = ""
	respondJSON(c, 200, "Profile detail", user)
}

// ProfileUpdateInput defines the parameters for updating a user's profile.
type ProfileUpdateInput struct {
	Nickname string `json:"nickname" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
	Info     string `json:"info"`
}

// UpdateProfile updates the authenticated user's profile.
func UpdateProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		respondJSON(c, 404, "User not found", nil)
		return
	}

	var input ProfileUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondJSON(c, 400, err.Error(), nil)
		return
	}

	user.Nickname = input.Nickname
	user.Sex = input.Sex
	user.Info = input.Info
	if err := db.Save(&user).Error; err != nil {
		respondJSON(c, 500, "Failed to update profile", nil)
		return
	}
	user.Password = "" // do not expose the password
	respondJSON(c, 200, "Profile updated", user)
}

// ----- Main ----- //

func main() {
	initDB()
	router := gin.Default()

	// Group endpoints under /usercenterr/v2.
	api := router.Group("/usercenter/v2")
	{
		// Registration and Login endpoints (POST only)
		api.POST("/register", Register)
		api.POST("/login", Login)

		// Profile endpoints (JWT protected)
		api.GET("/profile", JWTAuthMiddleware(), GetProfile)
		api.POST("/profile", JWTAuthMiddleware(), UpdateProfile)
	}

	// Run the server on port 8080.
	router.Run(":8080")
}
