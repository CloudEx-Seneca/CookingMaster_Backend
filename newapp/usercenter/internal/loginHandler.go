package internal

import "github.com/gin-gonic/gin"

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
