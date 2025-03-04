package recipe

import (
	"net/http"

	"CookingMaster_Backend/app/recipe/api/internal/logic/recipe"
	"CookingMaster_Backend/app/recipe/api/internal/svc"
	"CookingMaster_Backend/app/recipe/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RecipeDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecipeDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := recipe.NewRecipeDeleteLogic(r.Context(), svcCtx)
		resp, err := l.RecipeDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
