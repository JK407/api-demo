package handler

import (
	"net/http"

	"api-demo/internal/logic"
	"api-demo/internal/svc"
	"api-demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HelloWorldHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloworldReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHelloWorldLogic(r.Context(), svcCtx)
		resp, err := l.HelloWorld(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
