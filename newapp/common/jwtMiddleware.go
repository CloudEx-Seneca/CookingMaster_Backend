package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

// JWT secret – in production, store securely.
var JWT_SECRET = []byte("mysecretkey")

// JWTAuthMiddleware validates the token and sets the user_id in the context.
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			RespondJSON(c, 401, "Authorization header required", nil)
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			RespondJSON(c, 401, "Authorization header format must be Bearer {token}", nil)
			c.Abort()
			return
		}
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate signing method.
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return JWT_SECRET, nil
		})
		if err != nil || !token.Valid {
			RespondJSON(c, 401, "Invalid token", nil)
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			RespondJSON(c, 401, "Invalid token claims", nil)
			c.Abort()
			return
		}
		userID, ok := claims["user_id"].(float64)
		if !ok {
			RespondJSON(c, 401, "Invalid token user_id", nil)
			c.Abort()
			return
		}
		c.Set("user_id", uint(userID))
		c.Next()
	}
}
