package user

import (
	"CookingMaster_Backend/pkg/result"
	"net/http"

	"CookingMaster_Backend/app/usercenter/api/internal/logic/user"
	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewUserDetailLogic(r.Context(), svcCtx)
		resp, err := l.UserDetail(&req)
		result.HttpResult(r, w, resp, err)
	}
}
