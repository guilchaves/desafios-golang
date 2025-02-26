package main

import (
	"net/http"

	"github.com/guilchaves/desafios-golang/desafio_03/configs"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/database"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/webserver/handlers"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/webserver/routes"
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
	r := routes.NewRouter(clientHandler)

	http.ListenAndServe(":8080", r)
}
