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
	suite.Equal(uint(1), clientResult.ID)

	suite.Equal(client.Name, clientResult.Name)
	suite.Equal(client.Cpf, clientResult.Cpf)
	suite.Equal(client.Income, clientResult.Income)
	suite.True(client.BirthDate.Equal(clientResult.BirthDate), "Birth dates should match")
	suite.Equal(client.Children, clientResult.Children)

}
