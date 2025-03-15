package internal

import "github.com/gin-gonic/gin"

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
