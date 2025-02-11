package svc

import (
	"CookingMaster_Backend/app/usercenter/model"
	"CookingMaster_Backend/app/usercenter/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel  model.UsersModel
	TokenModel model.UserTokensModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:     c,
		UserModel:  model.NewUsersModel(sqlConn, c.Cache),
		TokenModel: model.NewUserTokensModel(sqlConn, c.Cache),
	}
}
