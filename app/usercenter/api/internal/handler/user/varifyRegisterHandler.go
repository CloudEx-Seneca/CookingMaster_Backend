package user

import (
	"net/http"

	"CookingMaster_Backend/app/usercenter/api/internal/logic/user"
	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func VarifyRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VarifyRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewVarifyRegisterLogic(r.Context(), svcCtx)
		resp, err := l.VarifyRegister(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
