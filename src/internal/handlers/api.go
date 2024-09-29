package handlers

import (
	"golearn/src/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(router *chi.Mux) {
	router.Use(chimiddle.StripSlashes)

	router.Route("/api", func(r chi.Router) {

		r.Route("/account", func(acc chi.Router) {
			acc.Use(middleware.Authorization)

			acc.Get("/balance", GetPointBalance)
		})
	})
}
