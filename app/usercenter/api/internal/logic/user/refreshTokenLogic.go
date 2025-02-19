package user

import (
	"CookingMaster_Backend/app/usercenter/model"
	"CookingMaster_Backend/pkg/authhelper"
	"CookingMaster_Backend/pkg/xerr"
	"context"
	"time"

	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenReq) (resp *types.RefreshTokenResp, err error) {
	rtm := authhelper.NewTokenParser(l.svcCtx.Config.JwtAuth.AccessSecret, req.RefreshToken)
	err = rtm.VarifyToken()
	if err != nil {
		return nil, xerr.NewCodeError(xerr.TOKEN_INVALID_ERROR)
	}

	now := time.Now().Unix()
	expire := now + l.svcCtx.Config.JwtAuth.AccessExpire
	atm := authhelper.NewTokenGenerator(l.svcCtx.Config.JwtAuth.AccessSecret, now, expire, rtm.GetUserId(), model.AccessTokenType)
	err = atm.GenerateJwtToken()
	if err != nil {
		return nil, err
	}
	refreshAfter := now + int64(float64(l.svcCtx.Config.JwtAuth.AccessExpire)*0.8)

	return &types.RefreshTokenResp{
		AccessToken:  atm.GetToken(),
		AccessExpire: expire,
		RefreshAfter: refreshAfter,
	}, nil
}
