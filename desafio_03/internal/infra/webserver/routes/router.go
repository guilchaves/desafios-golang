package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/webserver/handlers"
)

func NewRouter(clientHandler *handlers.ClientHandler) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/clients", func(r chi.Router) {
		r.Get("/", clientHandler.GetClients)
		r.Get("/{id}", clientHandler.GetClientByID)
		r.Post("/", clientHandler.CreateClient)
		r.Put("/{id}", clientHandler.UpdateClient)
		r.Delete("/{id}", clientHandler.DeleteProduct)
	})

	return r
}
