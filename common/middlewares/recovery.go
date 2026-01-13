package middlewares

import (
	"fmt"
	"net/http"

	"github.com/lxjianbai/swift/common/response"
)

func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if res := recover(); res != nil {
				// 触发脱敏机制，Response 会记录日志并返回 500
				response.HttpJson(r, w, nil, fmt.Errorf("[Panic] %v", res))
			}
		}()
		next(w, r)
	}
}
