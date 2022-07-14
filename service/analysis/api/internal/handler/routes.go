// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"ayzy_platform/service/analysis/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/analysis/getcustombyprovince/:proname",
				Handler: getcustombyprovinceHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/analysis/customsummary/:ddcode",
				Handler: customsummaryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/analysis/goodsqtysummary/:ddcode",
				Handler: goodsqtysummaryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/analysis/settle/summary",
				Handler: settlesummaryHandler(serverCtx),
			},
		},
	)
}