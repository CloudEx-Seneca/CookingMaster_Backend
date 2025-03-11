package svc

import (
	"CookingMaster_Backend/app/usercenter/api/internal/config"
	"CookingMaster_Backend/app/usercenter/model"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"os"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  model.UsersModel
	TokenModel model.UserTokensModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dataSource := c.DataSource
	if c.DBEnvEnabled {
		dbHost := os.Getenv("DB_HOST")
		mysqlPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
		dataSource = fmt.Sprintf("root:%s@tcp(%s:3306)/usercenter?charset=utf8mb4&parseTime=true&loc=Local", mysqlPassword, dbHost)
	}
	sqlConn := sqlx.NewMysql(dataSource)
	return &ServiceContext{
		Config:     c,
		UserModel:  model.NewUsersModel(sqlConn, c.Cache),
		TokenModel: model.NewUserTokensModel(sqlConn, c.Cache),
	}
}
