package logic

import (
	"CookingMaster_Backend/app/usercenter/model"
	"CookingMaster_Backend/pkg/ctxdata"
	"CookingMaster_Backend/pkg/xerr"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"

	"CookingMaster_Backend/app/usercenter/rpc/internal/svc"
	"CookingMaster_Backend/app/usercenter/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type VarifyTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrOneTimeTokenOverUsedError = xerr.NewCodeError(xerr.ONETIME_TOKEN_OVERUSED_ERROR)
var ErrTokenRevokedError = xerr.NewCodeError(xerr.TOKEN_REVOKED_ERROR)

func NewVarifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VarifyTokenLogic {
	return &VarifyTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VarifyTokenLogic) VarifyToken(in *usercenter.VarifyTokenReq) (*usercenter.VarifyTokenResp, error) {
	secretKey := l.svcCtx.Config.JwtAuth.AccessSecret
	claims := make(jwt.MapClaims)
	_, err := jwt.ParseWithClaims(in.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	userId := claims[ctxdata.CtxKeyJwtUserId].(int64)
	tokenType := claims[ctxdata.CtxKeyJwtTokenType].(int64)
	if tokenType == model.RegisterTokenType || tokenType == model.ResetTokenType {
		tokenInfo, _ := l.svcCtx.TokenModel.FindOneByUserIdType(l.ctx, userId, tokenType)
		if tokenInfo.Status == model.ActiveTokenStatus {
			tokenInfo.Status = model.UsedTokenStatus
			l.svcCtx.TokenModel.Update(l.ctx, tokenInfo)
		} else {
			return nil, errors.Wrapf(ErrOneTimeTokenOverUsedError, "VerifyToken err userId:%d, tokentype:%d",
				userId, tokenType)
		}
	} else if tokenType == model.RefreshTokenType {
		tokenInfo, _ := l.svcCtx.TokenModel.FindOneByUserIdType(l.ctx, userId, tokenType)
		if tokenInfo.Status == model.RevokedTokenStatus {
			return nil, errors.Wrapf(ErrTokenRevokedError, "VerifyToken err userId:%d, tokentype:%d",
				userId, tokenType)
		}
	}

	l.ctx = context.WithValue(l.ctx, ctxdata.CtxKeyJwtUserId, userId)
	l.ctx = context.WithValue(l.ctx, ctxdata.CtxKeyJwtTokenType, tokenType)
	return &usercenter.VarifyTokenResp{
		Code: 0,
	}, nil
}
