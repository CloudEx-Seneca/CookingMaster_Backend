package user

import (
	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"
	"CookingMaster_Backend/app/usercenter/model"
	"CookingMaster_Backend/app/usercenter/rpc/usercenterClient"
	"CookingMaster_Backend/pkg/ctxdata"
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
	//deadline := time.Now().Add(30 * time.Second)
	//ctx, cancel := context.WithDeadline(context.Background(), deadline)
	//defer cancel()
	//l.ctx = ctx
	_, err = l.svcCtx.UserCenterRpc.VarifyToken(l.ctx, &usercenterClient.VarifyTokenReq{
		Token: req.RegisterToken,
	})
	if err != nil {
		return nil, xerr.NewCodeError(xerr.TOKEN_INVALID_ERROR)
	}

	userId := ctxdata.GetUidFromCtx(l.ctx)
	user, _ := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	user.Status = model.VarifiedUserStatus
	l.svcCtx.UserModel.Update(l.ctx, user)

	token, _ := l.svcCtx.TokenModel.FindOneByUserIdType(l.ctx, userId, model.RegisterTokenType)
	token.Status = model.UsedTokenStatus
	l.svcCtx.TokenModel.Update(l.ctx, token)
	return &types.VarifyRegisterResp{}, nil
}
