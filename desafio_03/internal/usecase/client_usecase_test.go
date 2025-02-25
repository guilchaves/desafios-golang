package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockClientRepository struct {
	mock.Mock
}

var _ entity.ClientRepositoryInterface = (*MockClientRepository)(nil)

func (m *MockClientRepository) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *MockClientRepository) FindByID(id int) (*entity.Client, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *MockClientRepository) FindAll(page, limit int, sort string) ([]*entity.Client, error) {
	args := m.Called(page, limit, sort)
	return args.Get(0).([]*entity.Client), args.Error(1)
}

func (m *MockClientRepository) Update(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *MockClientRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

type ClientUseCaseTestSuite struct {
	suite.Suite
	mockRepo *MockClientRepository
	useCase  *usecase.ClientUseCase
	client   *entity.Client
}

func (suite *ClientUseCaseTestSuite) SetupTest() {
	suite.mockRepo = new(MockClientRepository)
	suite.useCase = usecase.NewClientUseCase(suite.mockRepo)
	suite.client = &entity.Client{
		ID:        1,
		Name:      "John Doe",
		Cpf:       "12345678901",
		Income:    5000.0,
		BirthDate: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		Children:  2,
	}
}

func (suite *ClientUseCaseTestSuite) TestCreateClient_WhenValidClient_ShouldSaveClient(
	t *testing.T,
) {
	suite.mockRepo.On("Save", suite.client).Return(nil)

	err := suite.useCase.Create(suite.client)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *ClientUseCaseTestSuite) TestFindAllClients_WhenClientsExist_ShouldReturnClients(
	t *testing.T,
) {
	clients := []*entity.Client{suite.client}
	suite.mockRepo.On("FindAll", 1, 10, "asc").Return(clients, nil)

	result, err := suite.useCase.GetClients(1, 10, "asc")

	suite.NoError(err)
	suite.Equal(clients, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *ClientUseCaseTestSuite) TestFindByID_WhenClientExists_ShouldReturnClient(
	t *testing.T,
) {
	suite.mockRepo.On("FindByID", suite.client.ID).Return(suite.client, nil)
	result, err := suite.useCase.GetClientByID(suite.client.ID)

	suite.NoError(err)
	suite.Equal(suite.client, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *ClientUseCaseTestSuite) TestFindByID_WhenClientNotExists_ShouldReturnError(
	t *testing.T,
) {
	suite.mockRepo.On("FindByID", 999).Return(nil, errors.New("client not found"))
	result, err := suite.useCase.GetClientByID(999)

	suite.Error(err)
	suite.Nil(result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *ClientUseCaseTestSuite) TestUpdateClient_WhenClientExists_ShouldSaveClient(
	t *testing.T,
) {
	suite.mockRepo.On("Update", suite.client).Return(nil)

	err := suite.useCase.Update(suite.client)
	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *ClientUseCaseTestSuite) TestUpdateClient_WhenClientNotExists_ShouldReturnError(
	t *testing.T,
) {
	suite.mockRepo.On("Update", suite.client).Return(errors.New("client not found"))

	err := suite.useCase.Update(suite.client)
	suite.Error(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *ClientUseCaseTestSuite) TestDeleteClient_WhenClientExists_ShouldDeleteClient(
	t *testing.T,
) {
	suite.mockRepo.On("Delete", suite.client.ID).Return(nil)

	err := suite.useCase.Delete(suite.client.ID)
	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *ClientUseCaseTestSuite) TestDeleteClient_WhenClientNotExists_ShouldReturnError(
	t *testing.T,
) {
	suite.mockRepo.On("Delete", 999).Return(errors.New("client not found"))

	err := suite.useCase.Delete(999)
	suite.Error(err)
	suite.mockRepo.AssertExpectations(suite.T())
}
