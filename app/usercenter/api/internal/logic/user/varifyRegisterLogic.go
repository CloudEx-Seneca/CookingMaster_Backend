package user

import (
	"context"

	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
