package recipe

import (
	"CookingMaster_Backend/pkg/result"
	"net/http"

	"CookingMaster_Backend/app/recipe/api/internal/logic/recipe"
	"CookingMaster_Backend/app/recipe/api/internal/svc"
	"CookingMaster_Backend/app/recipe/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RecipeDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecipeDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := recipe.NewRecipeDetailLogic(r.Context(), svcCtx)
		resp, err := l.RecipeDetail(&req)
		result.HttpResult(r, w, resp, err)
	}
}
