package recipe

import (
	"context"

	"CookingMaster_Backend/app/recipe/api/internal/svc"
	"CookingMaster_Backend/app/recipe/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecipeDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecipeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecipeDeleteLogic {
	return &RecipeDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecipeDeleteLogic) RecipeDelete(req *types.RecipeDeleteReq) (resp *types.RecipeDeleteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
