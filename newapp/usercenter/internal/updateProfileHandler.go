package internal

import (
	"CookingMaster_Backend/newapp/common"
	"github.com/gin-gonic/gin"
)

// ProfileUpdateInput defines the parameters for updating a user's profile.
type ProfileUpdateInput struct {
	Nickname string `json:"nickname" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
	Info     string `json:"info"`
}

// UpdateProfile updates the authenticated user's profile.
// @Summary Update user profile
// @Description Update profile details of the logged in user.
// @Tags usercenter
// @Accept json
// @Produce json
// @Param user body ProfileUpdateInput true "User profile update info"
// @Success 200 {object} common.Response
// @Security BearerAuth
// @Router /profile [post]
func UpdateProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		common.RespondJSON(c, 404, "User not found", nil)
		return
	}

	var input ProfileUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		common.RespondJSON(c, 400, err.Error(), nil)
		return
	}

	user.Nickname = input.Nickname
	user.Sex = input.Sex
	user.Info = input.Info
	if err := db.Save(&user).Error; err != nil {
		common.RespondJSON(c, 500, "Failed to update profile", nil)
		return
	}
	user.Password = "" // do not expose the password
	common.RespondJSON(c, 200, "Profile updated", user)
}
