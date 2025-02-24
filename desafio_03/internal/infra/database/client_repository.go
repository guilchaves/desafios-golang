package database

import (
	"database/sql"
	"fmt"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
)

type ClientRepository struct {
	Db *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{Db: db}
}

func (r *ClientRepository) Save(client *entity.Client) error {
	birthDateStr := client.BirthDate.Format("2006-01-02")
	stmt, err := r.Db.Prepare(
		"INSERT INTO clients (name, cpf, income, birthdate, children) VALUES (?, ?, ?, ?, ?)",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(client.Name, client.Cpf, client.Income, birthDateStr, client.Children)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get the last inserted ID: %w", err)
	}
	client.ID = uint(id)

	return nil
}
