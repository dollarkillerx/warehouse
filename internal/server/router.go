package server

import (
	"github.com/dollarkillerx/warehouse/internal/config"
	"github.com/dollarkillerx/warehouse/pkg/utils"
	"github.com/go-chi/chi/v5"

	"net/http"
)

func (s *Server) router() {
	s.chi.Route("/v1", func(r chi.Router) {
		// 用户认证Check
		r.Post("/auth", s.AuthCheck)
		// put object
		r.Post("/put_object", MiddlewareAuth(s.ApiPutObject))
		// get object
		r.Get("/get_object", MiddlewareAuth(s.ApiGetObject))
		// del object
		r.Post("/del_object", MiddlewareAuth(s.ApiDelete))
		// remove bucket
		r.Post("/remove_bucket", MiddlewareAuth(s.ApiRemoveBucket))
		// download
		r.Get("/download", s.Download)
	})
}

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
