package handler

import (
	"ayzy_platform/common/response"
	"ayzy_platform/service/analysis/api/internal/logic"
	"ayzy_platform/service/analysis/api/internal/svc"
	"ayzy_platform/service/analysis/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func hsinfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HsInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHsinfoLogic(r.Context(), ctx)
		resp, err := l.Hsinfo(&req)
		response.Response(w, resp, err)

	}
}
