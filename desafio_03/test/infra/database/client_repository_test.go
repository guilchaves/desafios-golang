package database

import (
	"testing"
	"time"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/database"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/guilchaves/desafios-golang/desafio_03/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestClientRepository_CreateAndFindById(t *testing.T) {
	db := utils.SetupTestDB(t)
	repo := database.NewClientRepository(db)

	client := &entity.Client{
		Name:      "John Doe",
		CPF:       "12345678901",
		Income:    5000.0,
		BirthDate: time.Date(1990, 02, 02, 0, 0, 0, 0, time.UTC),
		Children:  2,
	}

	err := repo.Create(client)
	assert.NoError(t, err)
	assert.NotZero(t, client.ID)

	foundClient, err := repo.FindByID(client.ID)
	assert.NoError(t, err)
	assert.Equal(t, client.Name, foundClient.Name)
	assert.Equal(t, client.CPF, foundClient.CPF)
}

func TestClientRepository_FindAll(t *testing.T) {
	db := utils.SetupTestDB(t)
	repo := database.NewClientRepository(db)

	clients := []*entity.Client{
		{Name: "Alice", CPF: "11111111111", Income: 3000, BirthDate: time.Now(), Children: 1},
		{Name: "Bob", CPF: "22222222222", Income: 4000, BirthDate: time.Now(), Children: 0},
	}

	for _, c := range clients {
		err := repo.Create(c)
		assert.NoError(t, err)
	}

	foundClients, err := repo.FindAll(10, 0)
	assert.NoError(t, err)
	assert.Len(t, foundClients, 2)
}

func TestClientRepository_Update(t *testing.T) {
	db := utils.SetupTestDB(t)
	repo := database.NewClientRepository(db)

	client := &entity.Client{
		Name:      "Charlie",
		CPF:       "33333333333",
		Income:    4500,
		BirthDate: time.Now(),
		Children:  1,
	}

	err := repo.Create(client)
	assert.NoError(t, err)

	client.Name = "Charlie Brown"
	err = repo.Update(client)
	assert.NoError(t, err)
	updatedClient, err := repo.FindByID(client.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Charlie Brown", updatedClient.Name)
}

func TestClientRepository_Delete(t *testing.T) {
	db := utils.SetupTestDB(t)
	repo := database.NewClientRepository(db)

	client := &entity.Client{
		Name:      "Charlie",
		CPF:       "33333333333",
		Income:    5000,
		BirthDate: time.Now(),
		Children:  2,
	}

	err := repo.Create(client)
	assert.NoError(t, err)

	err = repo.Delete(client.ID)
	assert.NoError(t, err)

	deletedClient, err := repo.FindByID(client.ID)
	assert.Nil(t, deletedClient)
	assert.Error(t, err)
}
