package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

const (
	UnkownSexType = 0
	MaleSexType   = 1
	FemaleSexType = 2
)

const (
	UnvarifiedUserStatus = 0
	VarifiedUserStatus   = 1
)

const (
	RegisterTokenType = 0
	ResetTokenType    = 1
	RefreshTokenType  = 2
	AccessTokenType   = 3
)

const (
	ActiveTokenStatus  = 0
	UsedTokenStatus    = 1
	ExpiredTokenStatus = 2
	RevokedTokenStatus = 3
)
