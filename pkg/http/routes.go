package httpserver

import (
	handler "meli-challenge-compliance/pkg/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//NewRoutes Declara los recursos que se van a utilizar en el servicio con sus respectivas funciones
func NewRoutes(services *handler.DriveHandler) http.Handler {
	r := chi.NewRouter()
	r.Route(
		"/search-in-doc",
		func(r chi.Router) {
			r.Get("/{id}", services.GetInvoiceFromUser) //GET /search-in-doc/{id}
		},
	)
	r.Route("/file", func(r chi.Router) {
		r.Post("/", services.PostFile) //POST /file
	})
	return r
}
