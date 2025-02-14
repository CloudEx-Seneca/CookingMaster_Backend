package user

import (
	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"
	"CookingMaster_Backend/app/usercenter/model"
	"CookingMaster_Backend/pkg/authhelper"
	"CookingMaster_Backend/pkg/xerr"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type VarifyRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVarifyRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VarifyRegisterLogic {
	return &VarifyRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VarifyRegisterLogic) VarifyRegister(req *types.VarifyRegisterReq) (resp *types.VarifyRegisterResp, err error) {
	tm := authhelper.NewTokenParser(l.svcCtx.Config.JwtAuth.AccessSecret, req.RegisterToken)
	err = tm.VarifyToken()
	if err != nil {
		return nil, xerr.NewCodeError(xerr.TOKEN_INVALID_ERROR)
	}

	user, _ := l.svcCtx.UserModel.FindOne(l.ctx, tm.GetUserId())
	user.Status = model.VarifiedUserStatus
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return nil, err
	}

	return &types.VarifyRegisterResp{}, nil
}
