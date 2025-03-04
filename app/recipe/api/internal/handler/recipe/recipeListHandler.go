package recipe

import (
	"net/http"

	"CookingMaster_Backend/app/recipe/api/internal/logic/recipe"
	"CookingMaster_Backend/app/recipe/api/internal/svc"
	"CookingMaster_Backend/app/recipe/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RecipeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecipeListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := recipe.NewRecipeListLogic(r.Context(), svcCtx)
		resp, err := l.RecipeList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
