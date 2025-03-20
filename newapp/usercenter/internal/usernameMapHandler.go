package internal

import (
	"CookingMaster_Backend/newapp/common"
	"github.com/gin-gonic/gin"
)

// UserNameMapInput defines the payload for retrieving user names by IDs.
// @Description Request payload for user name map API.
type UserNameMapInput struct {
	UserIDs []uint `json:"user_ids" binding:"required"`
}

// GetUserNameMap retrieves user names for a given list of user IDs.
// @Summary Get user name map
// @Description Given a list of user IDs, returns a map with each user ID as key and the corresponding user name as value.
// @Tags usercenter
// @Accept json
// @Produce json
// @Param data body UserNameMapInput true "List of user IDs"
// @Success 200 {object} common.Response
// @Router /namemap [post]
func GetUserNameMap(c *gin.Context) {
	var input UserNameMapInput
	if err := c.ShouldBindJSON(&input); err != nil {
		common.RespondJSON(c, 400, err.Error(), nil)
		return
	}

	// Query the User table for all given user IDs.
	var users []User
	if err := db.Where("id IN ?", input.UserIDs).Find(&users).Error; err != nil {
		common.RespondJSON(c, 500, "Database error", nil)
		return
	}

	// Construct a map with user ID as key and user name (nickname) as value.
	result := make(map[uint]string)
	for _, u := range users {
		result[u.ID] = u.Nickname
	}
	common.RespondJSON(c, 200, "Success", result)
}
