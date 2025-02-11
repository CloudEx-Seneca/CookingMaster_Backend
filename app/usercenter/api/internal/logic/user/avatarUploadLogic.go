package user

import (
	"context"

	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AvatarUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAvatarUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AvatarUploadLogic {
	return &AvatarUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AvatarUploadLogic) AvatarUpload(req *types.UploadAvatarReq) (resp *types.UploadAvatarResp, err error) {
	// todo: add your logic here and delete this line

	return &types.UploadAvatarResp{}, nil
}
