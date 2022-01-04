package server

import (
	"github.com/dollarkillerx/warehouse/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"net/http"
)

type Server struct {
	chi *chi.Mux
}

func New() *Server {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	// cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   config.GetCORSAllowedOrigins(),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	return &Server{chi: router}
}

func (s *Server) Run() error {
	s.router()

	return http.ListenAndServe(config.GetListenAddr(), s.chi)
}
