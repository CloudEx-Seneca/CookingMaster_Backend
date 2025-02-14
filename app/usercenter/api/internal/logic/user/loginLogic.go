package user

import (
	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"
	"CookingMaster_Backend/app/usercenter/model"
	"CookingMaster_Backend/pkg/authhelper"
	"CookingMaster_Backend/pkg/xerr"
	"context"
	"time"

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
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil {
		return nil, xerr.NewCodeError(xerr.EMAIL_UNREGISTERED_ERROR)
	}
	if user.Status == model.UnvarifiedUserStatus {
		return nil, xerr.NewCodeError(xerr.USER_NOT_VERIFIED_ERROR)
	}

	pm := authhelper.NewPasswordDecoder(req.Password, user.Password)
	ok, err := pm.VerifyPassword()
	if err != nil || ok == false {
		return nil, xerr.NewCodeError(xerr.USER_PASSWORD_ERROR)
	}

	now := time.Now().Unix()
	accessExpire := now + l.svcCtx.Config.JwtAuth.AccessExpire
	atm := authhelper.NewTokenGenerator(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, user.Id, model.AccessTokenType)
	err = atm.GenerateJwtToken()
	if err != nil {
		return nil, err
	}
	accessToken := atm.GetToken()

	refreshExpire := now + authhelper.REFRESH_TOKEN_EXPIRE
	rtm := authhelper.NewTokenGenerator(l.svcCtx.Config.JwtAuth.AccessSecret, now, refreshExpire, user.Id, model.RefreshTokenType)
	err = rtm.GenerateJwtToken()
	if err != nil {
		return nil, err
	}
	refreshToken := rtm.GetToken()
	refreshAfter := now + int64(float64(l.svcCtx.Config.JwtAuth.AccessExpire)*0.8)

	return &types.LoginResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AccessExpire: accessExpire,
		RefreshAfter: refreshAfter,
	}, nil
}
