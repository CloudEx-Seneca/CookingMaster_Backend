package svc

import (
	"CookingMaster_Backend/app/usercenter/api/internal/config"
	"CookingMaster_Backend/app/usercenter/model"
	"CookingMaster_Backend/app/usercenter/rpc/usercenterClient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserCenterRpc usercenterClient.Usercenter
	UserModel     model.UsersModel
	TokenModel    model.UserTokensModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:        c,
		UserCenterRpc: usercenterClient.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
		UserModel:     model.NewUsersModel(sqlConn, c.Cache),
		TokenModel:    model.NewUserTokensModel(sqlConn, c.Cache),
	}
}
