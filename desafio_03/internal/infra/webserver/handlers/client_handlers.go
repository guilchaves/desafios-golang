package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/dto"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/database"
)

type ClientHandler struct {
	ClientRepository database.ClientRepository
}

func NewClientHandler(db database.ClientRepository) *ClientHandler {
	return &ClientHandler{ClientRepository: db}
}

func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var client dto.CreateClientInputDTO

	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	birthDate, err := time.Parse("2006-01-02", client.BirthDate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	clientInput, err := entity.NewClient(
		client.Name,
		client.Cpf,
		client.Income,
		birthDate,
		client.Children,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.ClientRepository.Save(clientInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ClientHandler) GetClients(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}
	sort := r.URL.Query().Get("sort")

	clients, err := h.ClientRepository.FindAll(pageInt, limitInt, sort)
  if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
    return
  }
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)
}
