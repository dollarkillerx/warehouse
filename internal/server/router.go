package server

import (
	"github.com/dollarkillerx/warehouse/internal/config"
	"github.com/dollarkillerx/warehouse/pkg/utils"

	"net/http"
)

// AuthCheck 校验用户是否合法
func (s *Server) AuthCheck(w http.ResponseWriter, r *http.Request) {
	accessKey := r.Header.Get("AccessKey")
	secretKey := r.Header.Get("SecretKey")

	if config.GetAccessKey() != accessKey || config.GetSecretKey() != secretKey {
		http.Error(w, http.StatusText(401), 401)
		return
	}

	utils.String(w, 200, "success")
}
