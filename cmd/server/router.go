package main

import (
	"go-ticket/internal/config"
	"go-ticket/internal/handler"
	"go-ticket/internal/middleware"
	"net/http"
)

type RouteGroup struct {
	prefix      string
	middlewares []func(http.Handler) http.Handler
	mux         *http.ServeMux
}

func SetupRouter(userHandler *handler.UserHandler, authHandler *handler.AuthHandler, cfg *config.Config) http.Handler {
	mux := http.NewServeMux()

	api := NewGroup(mux, "/api")

	// public
	api.Handle("/login", Method(http.MethodPost, authHandler.Login))
	api.Handle("/register", Method(http.MethodPost, authHandler.Register))

	// authenticated
	auth := api.Group("", middleware.AuthMiddleware(cfg.JWTSecret))

	auth.Handle("/profile", Method(http.MethodGet, userHandler.GetProfile))
	auth.Handle("/profile/update", Method(http.MethodPut, userHandler.UpdateProfile))

	return mux
}

func NewGroup(mux *http.ServeMux, prefix string, m ...func(http.Handler) http.Handler) *RouteGroup {
	return &RouteGroup{
		prefix:      prefix,
		middlewares: m,
		mux:         mux,
	}
}

func (g *RouteGroup) Handle(path string, handler http.Handler) {
	finalHandler := handler
	for i := len(g.middlewares) - 1; i >= 0; i-- {
		finalHandler = g.middlewares[i](finalHandler)
	}

	g.mux.Handle(g.prefix+path, finalHandler)
}

func (g *RouteGroup) Group(prefix string, m ...func(http.Handler) http.Handler) *RouteGroup {
	return &RouteGroup{
		prefix:      g.prefix + prefix,
		middlewares: append(g.middlewares, m...),
		mux:         g.mux,
	}
}

func Method(method string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.Header().Set("Allow", method)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h(w, r)
	}
}
