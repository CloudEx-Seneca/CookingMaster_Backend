package svc

import (
	"CookingMaster_Backend/app/recipe/api/internal/config"
	"CookingMaster_Backend/app/recipe/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	RecipeModel      model.RecipesModel
	IngredientsModel model.IngredientsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:           c,
		RecipeModel:      model.NewRecipesModel(sqlConn, c.Cache),
		IngredientsModel: model.NewIngredientsModel(sqlConn, c.Cache),
	}
}
