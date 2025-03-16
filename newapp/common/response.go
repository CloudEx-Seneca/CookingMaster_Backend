package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response defines the professional JSON response format.
type Response struct {
	Code    int         `json:"code"`    // e.g., 200 for success, 400/500 for errors
	Message string      `json:"message"` // descriptive message
	Data    interface{} `json:"data"`    // result payload (if any)
}

// respondJSON sends a JSON response with the specified code, message, and data.
func RespondJSON(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
