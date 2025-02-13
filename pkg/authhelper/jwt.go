package authhelper

import (
	"CookingMaster_Backend/pkg/ctxdata"
	"github.com/golang-jwt/jwt/v4"
)

type TokenManager struct {
	secretKey string
	iat       int64
	exp       int64
	userId    int64
	tokenType int64
	token     string
}

func NewTokenGenerator(secretKey string, iat int64, exp int64, userId int64, tokenType int64) *TokenManager {
	return &TokenManager{
		secretKey: secretKey,
		iat:       iat,
		exp:       exp,
		userId:    userId,
		tokenType: tokenType,
	}
}

func (tm *TokenManager) GenerateJwtToken() error {
	claims := make(jwt.MapClaims)
	claims["iat"] = tm.iat
	claims["exp"] = tm.exp
	claims[ctxdata.CtxKeyJwtUserId] = tm.userId
	claims[ctxdata.CtxKeyJwtTokenType] = tm.tokenType
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(tm.secretKey))
	if err != nil {
		return err
	}

	tm.token = token
	return nil
}

func (tm *TokenManager) GetToken() string {
	return tm.token
}

func NewTokenParser(secret string, token string) *TokenManager {
	return &TokenManager{
		secretKey: secret,
		token:     token,
	}
}

func (tm *TokenManager) VarifyToken() error {
	claims := make(jwt.MapClaims)
	_, err := jwt.ParseWithClaims(tm.token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(tm.secretKey), nil
	})
	if err != nil {
		return err
	}

	tm.iat = int64(claims["iat"].(float64))
	tm.exp = int64(claims["exp"].(float64))
	tm.userId = int64(claims[ctxdata.CtxKeyJwtUserId].(float64))
	tm.tokenType = int64(claims[ctxdata.CtxKeyJwtTokenType].(float64))
	return nil
}

func (tm *TokenManager) GetUserId() int64 {
	return tm.userId
}

func (tm *TokenManager) GetTokenType() int64 {
	return tm.tokenType
}
