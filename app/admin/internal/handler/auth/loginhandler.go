// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package auth

import (
	"net/http"

	"github.com/lxjianbai/swift/app/admin/internal/logic/auth"
	"github.com/lxjianbai/swift/app/admin/internal/svc"
	"github.com/lxjianbai/swift/app/admin/internal/types"
	"github.com/lxjianbai/swift/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			response.HttpJson(r, w, nil, err)
			return
		}

		l := auth.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)

		response.HttpJson(r, w, resp, err)
	}
}
