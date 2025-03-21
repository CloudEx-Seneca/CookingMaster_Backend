package internal

import (
	"CookingMaster_Backend/newapp/common"
	"github.com/gin-gonic/gin"
)

// LoginInput defines the expected parameters for user login.
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Login handles user login.
// It verifies whether the email is registered, checks the password, and returns a JWT on success.
// @Summary User login
// @Description Authenticate user and return JWT token.
// @Tags usercenter
// @Accept json
// @Produce json
// @Param credentials body LoginInput true "Login credentials"
// @Success 200 {object} common.Response
// @Router /login [post]
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		common.RespondJSON(c, 400, err.Error(), nil)
		return
	}

	var user User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		common.RespondJSON(c, 400, "Email not registered", nil)
		return
	}

	if !checkPasswordHash(input.Password, user.Password) {
		common.RespondJSON(c, 400, "Password wrong", nil)
		return
	}

	// Generate the JWT token
	token, err := generateJWT(user.ID, user.Email, user.Nickname)
	if err != nil {
		common.RespondJSON(c, 500, "Failed to generate token", nil)
		return
	}

	// Respond with token and user ID
	common.RespondJSON(c, 200, "Login successful", gin.H{
		"token":  token,
		"user_id": user.ID,  // Add the user ID here
	})
}
