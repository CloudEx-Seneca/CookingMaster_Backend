package recipe

import (
	"context"

	"CookingMaster_Backend/app/recipe/api/internal/svc"
	"CookingMaster_Backend/app/recipe/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecipeDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecipeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecipeDetailLogic {
	return &RecipeDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecipeDetailLogic) RecipeDetail(req *types.RecipeDetailReq) (resp *types.RecipeDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
