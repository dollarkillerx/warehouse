package server

import (
	"github.com/dollarkillerx/warehouse/internal/config"

	"net/http"
)

// MiddlewareAuth 校验用户是否合法
func MiddlewareAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessKey := r.Header.Get("AccessKey")
		secretKey := r.Header.Get("SecretKey")

		if config.GetAccessKey() != accessKey || config.GetSecretKey() != secretKey {
			http.Error(w, http.StatusText(401), 401)
			return
		}
		next.ServeHTTP(w, r)
	})
}
