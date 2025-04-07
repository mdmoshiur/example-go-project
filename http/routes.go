package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/mdmoshiur/example-go/internal/middleware"
)

func RegisterRoutes(r *chi.Mux, h *Handler) {
	apiRouter := chi.NewRouter()
	r.Mount("/api", apiRouter)
	// add middleware to api sub router
	apiRouter.Use(middleware.Auth)
	apiRouter.Use(middleware.Gzip)

	r.Post("/api/v1/login", h.UserHandler.Login)
	r.Post("/api/v1/users", h.UserHandler.StoreOrUpdate)

	apiRouter.Route("/v1/users", func(r chi.Router) {
		r.Get("/me", h.UserHandler.Details)
		r.Put("/{id}", h.UserHandler.StoreOrUpdate)
		r.Post("/logout", h.UserHandler.Logout)
	})
}
