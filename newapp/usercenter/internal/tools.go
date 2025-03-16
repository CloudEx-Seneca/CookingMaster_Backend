package internal

import (
	"CookingMaster_Backend/newapp/common"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// hashPassword hashes the plaintext password using bcrypt.
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPasswordHash compares a plaintext password with its hashed version.
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// generateJWT generates a JWT token with user_id and email claims.
func generateJWT(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(720 * time.Hour).Unix(), // token expires in 720 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(common.JWT_SECRET)
}
