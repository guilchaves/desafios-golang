package database_test

import (
	"testing"
	"time"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/database"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ClientRepositoryTestSuite struct {
	suite.Suite
	Db *gorm.DB
}

func (suite *ClientRepositoryTestSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	suite.NoError(err)
	suite.Db = db

	db.AutoMigrate(&entity.Client{})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ClientRepositoryTestSuite))
}

func (suite *ClientRepositoryTestSuite) TestCreateClient_WhenSave_ThenShouldSaveClient() {
	defaultBirthdate := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
	client, err := entity.NewClient(
		"John Doe",
		"12345678901",
		5000.0,
		defaultBirthdate,
		2,
	)
	suite.NoError(err)

	repo := database.NewClientRepository(suite.Db)
	err = repo.Save(client)
	suite.NoError(err)

	suite.NotZero(client.ID, "Client ID should be set after saving")

	var clientResult entity.Client
	result := suite.Db.First(&clientResult, "id = ?", client.ID)
	suite.NoError(result.Error)
	suite.Equal(1, clientResult.ID)

	suite.Equal(client.Name, clientResult.Name)
	suite.Equal(client.Cpf, clientResult.Cpf)
	suite.Equal(client.Income, clientResult.Income)
	suite.True(client.BirthDate.Equal(clientResult.BirthDate), "Birth dates should match")
	suite.Equal(client.Children, clientResult.Children)
}

func (suite *ClientRepositoryTestSuite) TestFindByID_WhenClienExists_ThenShouldReturnClient() {
	defaultBirthdate := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
	client, err := entity.NewClient(
		"John Doe",
		"12345678901",
		5000.0,
		defaultBirthdate,
		2,
	)
	suite.NoError(err)

	repo := database.NewClientRepository(suite.Db)
	err = repo.Save(client)
	suite.NoError(err)

	foundClient, err := repo.FindByID(client.ID)
	suite.NoError(err)

	suite.Equal(client.ID, foundClient.ID)
	suite.Equal(client.Name, foundClient.Name)
	suite.Equal(client.Cpf, foundClient.Cpf)
	suite.Equal(client.Income, foundClient.Income)
	suite.True(client.BirthDate.Equal(foundClient.BirthDate),
		"Expected birth date %s, got %s",
		client.BirthDate.Format(time.RFC3339),
		foundClient.BirthDate.Format(time.RFC3339),
	)
	suite.Equal(client.Children, foundClient.Children)
}

func (suite *ClientRepositoryTestSuite) TestFindByID_WhenClientNotExists_ThenShouldReturnError() {
	id := 999
	repo := database.NewClientRepository(suite.Db)

	_, err := repo.FindByID(id)

	suite.Error(err)
	suite.Equal(gorm.ErrRecordNotFound, err,
		"Expected record not found error, got %v", err,
	)
}

func (suite *ClientRepositoryTestSuite) TestUpdateClient_WhenClientExists_ThenShouldSaveClient() {
	defaultBirthdate := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
	client, err := entity.NewClient(
		"John Doe",
		"12345678901",
		5000.0,
		defaultBirthdate,
		2,
	)
	suite.NoError(err)

	repo := database.NewClientRepository(suite.Db)
	err = repo.Save(client)
	suite.NoError(err)

	client.Name = "John Updated"
	client.Cpf = "01987654321"
	err = repo.Update(client)
	suite.NoError(err)
	suite.Equal("John Updated", client.Name)
	suite.Equal("01987654321", client.Cpf)
}

func (suite *ClientRepositoryTestSuite) TestUpdateClient_WhenClientNotExists_ThenShouldReturnError() {
defaultBirthdate := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
	client, err := entity.NewClient(
		"John Doe",
		"12345678901",
		5000.0,
		defaultBirthdate,
		2,
	)
	suite.NoError(err)

	repo := database.NewClientRepository(suite.Db)
	client.Name = "John Updated"
	client.Cpf = "01987654321"
	err = repo.Update(client)
	suite.Error(err)
}
