package user

import (
	"CookingMaster_Backend/pkg/authhelper"
	"CookingMaster_Backend/pkg/xerr"
	"context"

	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutReq) (resp *types.LogoutResp, err error) {
	tm := authhelper.NewTokenParser(l.svcCtx.Config.JwtAuth.AccessSecret, req.RefreshToken)
	err = tm.VarifyToken()
	if err != nil {
		return nil, xerr.NewCodeError(xerr.TOKEN_INVALID_ERROR)
	}

	// TODO: invalid fresh token, for now leave it to frontend because jwt is stateless
	return &types.LogoutResp{}, nil
}
