package handler

import (
	"net/http"

	"github.com/jianxinliu/docker-starter/app/internal/logic"
	"github.com/jianxinliu/docker-starter/app/internal/svc"
	"github.com/jianxinliu/docker-starter/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AppHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAppLogic(r.Context(), svcCtx)
		resp, err := l.App(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
