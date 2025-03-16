package internal

import (
	"CookingMaster_Backend/newapp/common"
	"github.com/gin-gonic/gin"
)

// GetProfile returns the profile details of the authenticated user.
func GetProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		common.RespondJSON(c, 404, "User not found", nil)
		return
	}
	// Exclude the password field.
	user.Password = ""
	common.RespondJSON(c, 200, "Profile detail", user)
}
