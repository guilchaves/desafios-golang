package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/dto"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/database"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/webserver/response"
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
		slog.Error(err.Error())
		response.SendJSON(w, response.Response{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	birthDate, err := time.Parse("2006-01-02", client.BirthDate)
	if err != nil {
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusUnprocessableEntity, err.Error()),
			http.StatusUnprocessableEntity,
		)
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
		slog.Error(err.Error())
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusUnprocessableEntity, err.Error()),
			http.StatusUnprocessableEntity,
		)
		return
	}

	err = h.ClientRepository.Save(clientInput)
	if err != nil {
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusInternalServerError, err.Error()),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
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
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusInternalServerError, err.Error()),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)
}

func (h *ClientHandler) GetClientByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusBadRequest, ""),
			http.StatusBadRequest,
		)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusInternalServerError, err.Error()),
			http.StatusInternalServerError,
		)
		return
	}

	client, err := h.ClientRepository.FindByID(idInt)
	if err != nil {
		slog.Error(err.Error())
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusNotFound, err.Error()),
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

func (h *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusBadRequest, ""),
			http.StatusBadRequest,
		)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusInternalServerError, err.Error()),
			http.StatusInternalServerError,
		)
		return
	}

	var client dto.UpdateClientInputDTO
	err = json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		slog.Error(err.Error())
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusBadRequest, err.Error()),
			http.StatusBadRequest,
		)
		return
	}

	birthDate, err := time.Parse("2006-01-02", client.BirthDate)
	if err != nil {
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusBadRequest, err.Error()),
			http.StatusBadRequest,
		)
		return
	}

	_, err = h.ClientRepository.FindByID(idInt)
	if err != nil {
		slog.Error(err.Error())
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusNotFound, err.Error()),
			http.StatusNotFound,
		)
		return
	}

	clientInput := entity.Client{
		ID:        idInt,
		Name:      client.Name,
		Cpf:       client.Cpf,
		Income:    client.Income,
		BirthDate: birthDate,
		Children:  client.Children,
	}

	err = h.ClientRepository.Update(&clientInput)
	if err != nil {
		slog.Error(err.Error())
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusNotFound, err.Error()),
			http.StatusNotFound,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ClientHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusBadRequest, ""),
			http.StatusBadRequest,
		)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusInternalServerError, err.Error()),
			http.StatusInternalServerError,
		)

		return
	}

	_, err = h.ClientRepository.FindByID(idInt)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusNotFound, err.Error()),
			http.StatusNotFound,
		)
		return
	}

	err = h.ClientRepository.Delete(idInt)
	if err != nil {
		response.SendJSON(
			w,
			*response.ErrorResponse(http.StatusInternalServerError, err.Error()),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}
