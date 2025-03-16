package internal

import (
	"CookingMaster_Backend/newapp/common"
	"github.com/gin-gonic/gin"
)

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
		common.RespondJSON(c, 400, err.Error(), nil)
		return
	}

	// Check if the email is already registered.
	var user User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err == nil {
		common.RespondJSON(c, 400, "Email already registered", nil)
		return
	}

	// Hash password and create user.
	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		common.RespondJSON(c, 500, "Failed to hash password", nil)
		return
	}
	user = User{
		Email:    input.Email,
		Password: hashedPassword,
	}
	if err := db.Create(&user).Error; err != nil {
		common.RespondJSON(c, 500, "Failed to create user", nil)
		return
	}
	common.RespondJSON(c, 200, "Register successful", nil)
}
