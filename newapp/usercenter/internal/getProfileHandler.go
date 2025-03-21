package internal

import (
	"CookingMaster_Backend/newapp/common"
	"github.com/gin-gonic/gin"
)

// GetProfile returns the profile details of the authenticated user.
// @Summary Get user profile
// @Description Retrieve profile details of the logged in user.
// @Tags usercenter
// @Produce json
// @Success 200 {object} common.Response
// @Security BearerAuth
// @Router /profile [get]
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
