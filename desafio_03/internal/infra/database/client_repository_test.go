package database_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/database"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type ClientRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *ClientRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec(`
    CREATE TABLE clients (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        cpf TEXT NOT NULL UNIQUE,
        income REAL NOT NULL,
        birthdate TEXT NOT NULL,
        children INTEGER NOT NULL
    )
  `)
	suite.Db = db
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ClientRepositoryTestSuite))
}

func (suite *ClientRepositoryTestSuite) TestCreateClient_WhenSave_ThenShouldSaveClient() {
	var (
		id           uint
		name         string
		cpf          string
		income       float64
		birthDateStr string
		children     uint
	)

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

	var insertedID int64
	err = suite.Db.QueryRow("SELECT last_insert_rowid()").Scan(&insertedID)
	suite.NoError(err)
	suite.NotZero(insertedID, "Should have generated an ID")

	err = suite.Db.QueryRow("SELECT id, name, cpf, income, birthdate, children FROM clients WHERE id = ?", insertedID).
		Scan(&id, &name, &cpf, &income, &birthDateStr, &children)
	suite.NoError(err)

	birthDate, err := time.Parse("2006-01-02", birthDateStr)
	suite.NoError(err)

	clientResult := entity.Client{
		ID:        id,
		Name:      name,
		Cpf:       cpf,
		Income:    income,
		BirthDate: birthDate,
		Children:  children,
	}

	suite.Equal(client.ID, clientResult.ID)
	suite.Equal(client.Name, clientResult.Name)
	suite.Equal(client.Cpf, clientResult.Cpf)
	suite.Equal(client.Income, clientResult.Income)
	suite.Equal(client.BirthDate, clientResult.BirthDate)
	suite.Equal(client.Children, clientResult.Children)
}
