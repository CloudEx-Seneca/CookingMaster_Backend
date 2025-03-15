package internal

import "github.com/gin-gonic/gin"

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
