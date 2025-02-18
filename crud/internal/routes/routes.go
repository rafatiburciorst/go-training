package routes

import (
	"crud/internal/handlers"
	"crud/internal/models"

	"github.com/go-chi/chi/v5"
)

func Init(db map[int64]models.User) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.ListUsers(db))
	})

	return router
}
