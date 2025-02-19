package user

import (
	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"
	"CookingMaster_Backend/pkg/ctxdata"
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateReq) (resp *types.UserUpdateResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}

	user.Nickname = sql.NullString{String: req.Nickname, Valid: req.Nickname != ""}
	user.Sex = req.Sex
	user.Info = sql.NullString{String: req.Info, Valid: req.Info != ""}
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return nil, err
	}

	return &types.UserUpdateResp{}, nil
}
