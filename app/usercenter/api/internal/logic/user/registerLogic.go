package user

import (
	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"
	"CookingMaster_Backend/app/usercenter/model"
	"CookingMaster_Backend/pkg/authhelper"
	"CookingMaster_Backend/pkg/email"
	"CookingMaster_Backend/pkg/xerr"
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if user != nil {
		return nil, xerr.NewCodeError(xerr.EMAIL_REGISTERD_BEFORE_ERROR)
	}

	userId, err := authhelper.GenerateUserId()
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	expire := now + l.svcCtx.Config.JwtAuth.AccessExpire
	tg := authhelper.NewTokenGenerator(l.svcCtx.Config.JwtAuth.AccessSecret, now, expire, userId, model.RegisterTokenType)
	err = tg.GenerateJwtToken()
	if err != nil {
		return nil, err
	}
	token := tg.GetToken()

	subject := email.REGISTER_EMAIL_SUBJECT
	body := fmt.Sprintf(email.REGISTER_EMAIL_BODY_TEMPLATE, token)
	go email.SendEmail(req.Email, subject, body)

	pm := authhelper.NewPasswordEncoder(req.Password)
	err = pm.EncodedHash()
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.Users{
		Id:       userId,
		Email:    req.Email,
		Password: pm.GetEncodedHash(),
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		RegisterToken:  token,
		RegisterExpire: expire,
	}, nil
}
