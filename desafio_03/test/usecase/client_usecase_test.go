package usecase

import (
	"testing"
	"time"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClientRepository struct {
	mock.Mock
}

func setupClientUsecase(t *testing.T) (*MockClientRepository, *usecase.ClientUsecase) {
	t.Helper()
	mockRepo := new(MockClientRepository)
	clientUsecase := usecase.NewClientUsecase(mockRepo)
	return mockRepo, clientUsecase
}

func (m *MockClientRepository) Create(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *MockClientRepository) FindByID(id uint) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *MockClientRepository) FindAll(limit, offset int) ([]entity.Client, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]entity.Client), args.Error(1)
}

func (m *MockClientRepository) Update(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *MockClientRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestClientUsecase_CreateClient(t *testing.T) {
	mockRepo, clientUsecase := setupClientUsecase(t)

	client := &entity.Client{
		Name:      "John Doe",
		CPF:       "12345678901",
		Income:    5000.0,
		BirthDate: time.Now().UTC().Truncate(time.Second),
		Children:  2,
	}

	mockRepo.On("Create", client).Return(nil)
	mockRepo.On("Create", mock.AnythingOfType("*entity.Client")).Return(nil)

	createdClient, err := clientUsecase.CreateClient(
		client.Name, client.CPF, client.Income,
		client.BirthDate.Format("2006-01-02"), client.Children,
	)

	assert.NoError(t, err)
	assert.NotNil(t, createdClient)
	mockRepo.AssertCalled(t, "Create", mock.AnythingOfType("*entity.Client"))
}

func TestClientUsecase_GetClientByID(t *testing.T) {
	mockRepo, clientUsecase := setupClientUsecase(t)

	client := &entity.Client{
		ID:        1,
		Name:      "Alice",
		CPF:       "22222222222",
		Income:    4500,
		BirthDate: time.Now().UTC().Truncate(time.Second),
		Children:  1,
	}

	mockRepo.On("FindByID", client.ID).Return(client, nil)

	foundClient, err := clientUsecase.GetClientByID(client.ID)

	assert.NoError(t, err)
	assert.Equal(t, client.Name, foundClient.Name)
	mockRepo.AssertCalled(t, "FindByID", client.ID)
}

func TestClientUsecase_GetClients(t *testing.T) {
	mockRepo, clientUsecase := setupClientUsecase(t)

	clients := []entity.Client{
		{
			Name:      "Alice",
			CPF:       "11111111111",
			Income:    3000,
			BirthDate: time.Now().UTC().Truncate(time.Second),
			Children:  1,
		},
		{
			Name:      "Bob",
			CPF:       "22222222222",
			Income:    4000,
			BirthDate: time.Now().UTC().Truncate(time.Second),
			Children:  0,
		},
	}

	mockRepo.On("FindAll", 10, 0).Return(clients, nil)

	foundClients, err := clientUsecase.GetClients()

	assert.NoError(t, err)
	assert.Len(t, foundClients, 2)
	mockRepo.AssertCalled(t, "FindAll", 10, 0)
}

func TestClientUsecase_UpdateClient(t *testing.T) {
	mockRepo, clientUsecase := setupClientUsecase(t)

	client := &entity.Client{
		ID:        1,
		Name:      "Charlie",
		CPF:       "33333333333",
		Income:    4500,
		BirthDate: time.Now().UTC().Truncate(time.Second),
		Children:  1,
	}

	mockRepo.On("Update", client).Return(nil)

	err := clientUsecase.UpdateClient(client)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Update", client)
}

func TestClientUsecase_DeleteClient(t *testing.T) {
	mockRepo, clientUsecase := setupClientUsecase(t)

	clientID := uint(1)
	mockRepo.On("Delete", clientID).Return(nil)

	err := clientUsecase.DeleteClient(clientID)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Delete", clientID)
}
