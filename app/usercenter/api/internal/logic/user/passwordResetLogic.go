package user

import (
	"CookingMaster_Backend/app/usercenter/model"
	"CookingMaster_Backend/pkg/authhelper"
	"CookingMaster_Backend/pkg/email"
	"CookingMaster_Backend/pkg/xerr"
	"context"
	"fmt"
	"time"

	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PasswordResetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPasswordResetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PasswordResetLogic {
	return &PasswordResetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PasswordResetLogic) PasswordReset(req *types.PasswordResetReq) (resp *types.PasswordResetResp, err error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil {
		return nil, xerr.NewCodeError(xerr.EMAIL_UNREGISTERED_ERROR)
	}

	now := time.Now().Unix()
	expire := now + l.svcCtx.Config.JwtAuth.AccessExpire
	tm := authhelper.NewTokenGenerator(l.svcCtx.Config.JwtAuth.AccessSecret, now, expire, user.Id, model.ResetTokenType)
	err = tm.GenerateJwtToken()
	if err != nil {
		return nil, err
	}

	body := fmt.Sprintf(email.RESET_EMAIL_BODY_TEMPLATE, l.svcCtx.Config.EmailLinkDomain, tm.GetToken())
	go email.SendEmail(req.Email, email.RESET_EMAIL_SUBJECT, body)

	return &types.PasswordResetResp{
		ResetToken:  tm.GetToken(),
		ResetExipre: expire,
	}, nil
}
