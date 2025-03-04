package recipe

import (
	"CookingMaster_Backend/app/recipe/api/internal/svc"
	"CookingMaster_Backend/app/recipe/api/internal/types"
	"CookingMaster_Backend/app/recipe/model"
	"CookingMaster_Backend/pkg/ctxdata"
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecipeInsertOrUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecipeInsertOrUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecipeInsertOrUpdateLogic {
	return &RecipeInsertOrUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecipeInsertOrUpdateLogic) RecipeInsertOrUpdate(req *types.RecipeInsertoOrUpdateReq) (resp *types.RecipeInsertOrUpdateResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	recipe, err := l.svcCtx.RecipeModel.FindOneByName(l.ctx, req.Recipe.Name)
	if err != nil {
		_, err := l.svcCtx.RecipeModel.Insert(l.ctx, &model.Recipes{
			UserId:      userId,
			Name:        req.Recipe.Name,
			Description: sql.NullString{String: req.Recipe.Description, Valid: req.Recipe.Description != ""},
		})
		//recipeId, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}

		//for i := 0; i < len(req.Recipe.Ingredients); i++ {
		//	_, err = l.svcCtx.IngredientsModel.Insert(l.ctx, &model.Ingredients{
		//		RecipeId: recipeId,
		//		Name:     req.Recipe.Ingredients[i],
		//		Quantity: sql.NullString{String: req.Recipe.Quantities[i], Valid: req.Recipe.Quantities[i] != ""},
		//	})
		//}
		//if err != nil {
		//	return nil, err
		//}
		return &types.RecipeInsertOrUpdateResp{}, nil
	}

	recipe.Description = sql.NullString{String: req.Recipe.Description, Valid: req.Recipe.Description != ""}
	err = l.svcCtx.RecipeModel.Update(l.ctx, recipe)
	if err != nil {
		return nil, err
	}

	return &types.RecipeInsertOrUpdateResp{}, nil
}
