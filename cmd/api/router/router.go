package router

import (
	"myapp/cmd/resource/book"
	"myapp/cmd/resource/health"

	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/livez", health.Read)

	r.Route("/v1", func(r chi.Router) {
		bookApi := &book.API{}

		r.Get("/books", bookApi.List)
		r.Post("/books", bookApi.Create)
		r.Get("/books/{id}", bookApi.Read)
		r.Put("/books/{id}", bookApi.Update)
		r.Delete("/books/{id}", bookApi.Delete)
	})

	return r
}
