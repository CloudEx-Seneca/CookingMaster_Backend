package user

import (
	"context"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"

	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return &types.LoginResp{}, nil
}

func VerifyPassword(password string, encodedHash string) (bool, error) {
	var algo string
	var time, memory uint32
	var threads uint8
	var saltB64, hashB64 string

	_, err := fmt.Scanf(encodedHash, "%s$%d$%d$%d$%s$%s", &algo, &time, &memory, &threads, &saltB64, &hashB64)
	if err != nil {
		return false, err
	}

	salt, err := base64.StdEncoding.DecodeString(saltB64)
	if err != nil {
		return false, err
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(hashB64)
	if err != nil {
		return false, err
	}

	newHash := argon2.IDKey([]byte(password), salt, time, memory, threads, uint32(len(expectedHash)))
	return string(newHash) == string(expectedHash), nil
}
