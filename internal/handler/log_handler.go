package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lottery/internal/logic"
	"lottery/internal/svc"
	"lottery/internal/types"
)

// Pick a team by encryptCode
func LogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLogLogic(r.Context(), svcCtx)
		resp, err := l.Log(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}