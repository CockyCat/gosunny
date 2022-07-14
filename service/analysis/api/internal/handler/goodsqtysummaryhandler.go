package handler

import (
	"net/http"

	"ayzy_platform/service/analysis/api/internal/logic"
	"ayzy_platform/service/analysis/api/internal/svc"
	"ayzy_platform/service/analysis/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func goodsqtysummaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGoodsqtysummaryLogic(r.Context(), svcCtx)
		resp, err := l.Goodsqtysummary(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
