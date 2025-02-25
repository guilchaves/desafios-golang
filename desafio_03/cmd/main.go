package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/guilchaves/desafios-golang/desafio_03/configs"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/database"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Client{})
	clientRepo := database.NewClientRepository(db)
	clientHandler := handlers.NewClientHandler(*clientRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/clients", clientHandler.GetClients)
	r.Get("/clients/{id}", clientHandler.GetClientByID)
	r.Post("/clients", clientHandler.CreateClient)
	r.Put("/clients/{id}", clientHandler.UpdateClient)
	r.Delete("/clients/{id}", clientHandler.DeleteProduct)

	http.ListenAndServe(":8080", r)
}
