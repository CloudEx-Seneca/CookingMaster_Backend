package user

import (
	"CookingMaster_Backend/pkg/authhelper"
	"CookingMaster_Backend/pkg/xerr"
	"context"

	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordLogic) ResetPassword(req *types.ResetPasswordReq) (resp *types.ResetPasswordResp, err error) {
	tm := authhelper.NewTokenParser(l.svcCtx.Config.JwtAuth.AccessSecret, req.ResetToken)
	err = tm.VarifyToken()
	if err != nil {
		return nil, xerr.NewCodeError(xerr.TOKEN_INVALID_ERROR)
	}

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, tm.GetUserId())
	if err != nil {
		return nil, err
	}
	pm := authhelper.NewPasswordEncoder(req.Password)
	err = pm.EncodedHash()
	if err != nil {
		return nil, err
	}
	user.Password = pm.GetEncodedHash()
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return nil, err
	}

	return &types.ResetPasswordResp{}, nil
}
