package recipe

import (
	"CookingMaster_Backend/pkg/result"
	"net/http"

	"CookingMaster_Backend/app/recipe/api/internal/logic/recipe"
	"CookingMaster_Backend/app/recipe/api/internal/svc"
	"CookingMaster_Backend/app/recipe/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RecipeInsertOrUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecipeInsertoOrUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := recipe.NewRecipeInsertOrUpdateLogic(r.Context(), svcCtx)
		resp, err := l.RecipeInsertOrUpdate(&req)
		result.HttpResult(r, w, resp, err)
	}
}
