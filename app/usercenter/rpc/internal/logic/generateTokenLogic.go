package logic

import (
	"CookingMaster_Backend/app/usercenter/model"
	"CookingMaster_Backend/pkg/ctxdata"
	"CookingMaster_Backend/pkg/xerr"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"

	"CookingMaster_Backend/app/usercenter/rpc/internal/svc"
	"CookingMaster_Backend/app/usercenter/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrGenerateTokenError = xerr.NewCodeError(xerr.TOKEN_GENERATE_ERROR)

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *usercenter.GenerateTokenReq) (*usercenter.GenerateTokenResp, error) {
	now := time.Now().Unix()
	secretKey := l.svcCtx.Config.JwtAuth.AccessSecret
	var token string
	var expire int64
	var err error
	if in.TokenType == model.RegisterTokenType || in.TokenType == model.ResetTokenType || in.TokenType == model.AccessTokenType {
		expire = l.svcCtx.Config.JwtAuth.ShortExpire
		token, err = l.generateJwtToken(secretKey, now, expire, in.UserId, in.TokenType)
	} else if in.TokenType == model.RefreshTokenType {
		expire = l.svcCtx.Config.JwtAuth.LongExpire
		token, err = l.generateJwtToken(secretKey, now, expire, in.UserId, in.TokenType)
	}
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "generateJwtToken err userId:%d, tokentype:%d, err:%v",
			in.UserId, in.TokenType, err)
	}

	if in.TokenType == model.RegisterTokenType || in.TokenType == model.ResetTokenType || in.TokenType == model.RefreshTokenType {
		l.svcCtx.TokenModel.Update(l.ctx, &model.UserTokens{
			UserId: in.UserId,
			Type:   in.TokenType,
			Token:  token,
			Expire: time.Unix(now+expire, 0),
			Status: model.ActiveTokenStatus,
		})
	}

	return &usercenter.GenerateTokenResp{
		Token:  token,
		Expire: now + expire,
	}, nil
}

func (l *GenerateTokenLogic) generateJwtToken(secretKey string, iat, seconds, userId, tokenType int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["iat"] = iat
	claims["exp"] = iat + seconds
	claims[ctxdata.CtxKeyJwtUserId] = userId
	claims[ctxdata.CtxKeyJwtTokenType] = tokenType
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
}
