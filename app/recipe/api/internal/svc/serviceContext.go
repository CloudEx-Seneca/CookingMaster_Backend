package svc

import (
	"CookingMaster_Backend/app/recipe/api/internal/config"
	"CookingMaster_Backend/app/recipe/model"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"os"
)

type ServiceContext struct {
	Config config.Config

	RecipeModel      model.RecipesModel
	IngredientsModel model.IngredientsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dataSource := c.DataSource
	if c.DBEnvEnabled {
		dbHost := os.Getenv("DB_HOST")
		mysqlPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
		dataSource = fmt.Sprintf("root:%s@tcp(%s:3306)/recipe?charset=utf8mb4&parseTime=true&loc=Local", mysqlPassword, dbHost)
	}
	sqlConn := sqlx.NewMysql(dataSource)
	return &ServiceContext{
		Config:           c,
		RecipeModel:      model.NewRecipesModel(sqlConn, c.Cache),
		IngredientsModel: model.NewIngredientsModel(sqlConn, c.Cache),
	}
}
