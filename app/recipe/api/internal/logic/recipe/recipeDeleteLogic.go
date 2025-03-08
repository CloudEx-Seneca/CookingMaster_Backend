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
	// params := mux.Vars(r)
	// id, err := strconv.Atoi(params["id"])
	// if err != nil {
		// http.Error(w, "Invalid recipe ID", http.StatusBadRequest)
		// return
	// }

	// for i, recipe := range recipes {
		// if recipe.ID == id {
			// recipes = append(recipes[:i], recipes[i+1:]...)
			// w.WriteHeader(http.StatusNoContent)
			// return
		// }
	// }

	// http.Error(w, "Recipe not found", http.StatusNotFound)
// }

	// return
// }
