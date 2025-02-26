package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/guilchaves/desafios-golang/desafio_03/configs"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/database"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/webserver/handlers"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/webserver/routes"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("Starting application...")

	_, err := configs.LoadConfig(".")
	if err != nil {
		slog.Error("failed to load config", "error", err)
		panic(err)
	}

	db := database.NewDB()
	db.AutoMigrate(&entity.Client{})

	clientRepo := database.NewClientRepository(db)
	clientHandler := handlers.NewClientHandler(*clientRepo)

	r := routes.NewRouter(clientHandler)

	slog.Info("Server started", "port", 8080)
	http.ListenAndServe(":8080", r)
}
