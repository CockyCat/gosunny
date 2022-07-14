package handler

import (
	"net/http"

	"ayzy_platform/service/analysis/api/internal/logic"
	"ayzy_platform/service/analysis/api/internal/svc"
	"ayzy_platform/service/analysis/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getcustombyprovinceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProvinceCustomReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetcustombyprovinceLogic(r.Context(), svcCtx)
		resp, err := l.Getcustombyprovince(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
