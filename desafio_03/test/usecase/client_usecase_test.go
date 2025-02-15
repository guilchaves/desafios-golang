package usecase

import (
	"testing"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/repository"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/usecase"
	"github.com/guilchaves/desafios-golang/desafio_03/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestClientUsecase_Create(t *testing.T) {
	db := utils.SetupTestDB(t)

	clientRepo := repository.NewClientRepository(db)
	clientUsecase := usecase.NewClientUsecase(clientRepo)

	name := "John Doe"
	cpf := "12345678901"
	income := 5000.0
	birthDateStr := "1990-02-30"
	children := 2

	client, err := clientUsecase.Create(name, cpf, income, birthDateStr, children)
	assert.Error(t, err)
	assert.Nil(t, client)
	assert.Equal(t, "data de nascimento inválida", err.Error())
}
