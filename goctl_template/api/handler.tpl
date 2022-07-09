package {{.PkgName}}                                                                                                                   

import (
        "net/http"
        "github.com/zeromicro/go-zero/rest/httpx"
        "ayzy_platform/common/response"
        {{.ImportPackages}}
)

func {{.HandlerName}}(ctx *svc.ServiceContext) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
                {{if .HasRequest}}var req types.{{.RequestType}}
                if err := httpx.Parse(r, &req); err != nil {
                        httpx.Error(w, err)
                                return
                }{{end}}

                l := logic.New{{.LogicType}}(r.Context(), ctx)
                {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
                {{if .HasResp}}response.Response(w, resp, err){{else}}response.Response(w, nil, err){{end}}

        }   
}
