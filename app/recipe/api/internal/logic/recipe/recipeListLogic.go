package recipe

import (
	"context"

	"CookingMaster_Backend/app/recipe/api/internal/svc"
	"CookingMaster_Backend/app/recipe/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecipeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecipeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecipeListLogic {
	return &RecipeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecipeListLogic) RecipeList(req *types.RecipeListReq) (resp *types.RecipeListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
