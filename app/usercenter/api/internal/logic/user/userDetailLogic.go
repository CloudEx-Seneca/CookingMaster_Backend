package user

import (
	"CookingMaster_Backend/pkg/ctxdata"
	"context"

	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailReq) (resp *types.UserDetailResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}

	var userInfo = types.User{
		UserId:    user.Id,
		Email:     user.Email,
		Status:    user.Status,
		Nickname:  user.Nickname.String,
		Sex:       user.Sex,
		AvatarURL: user.AvatarUrl.String,
		Info:      user.Info.String,
	}

	return &types.UserDetailResp{
		UserDetail: userInfo,
	}, nil
}
