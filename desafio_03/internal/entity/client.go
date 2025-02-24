package entity

import (
	"errors"
	"time"
)

type Client struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Income    float64   `json:"income"`
	BirthDate time.Time `json:"birthDate"`
	Children  uint      `json:"children"`
}

func NewClient(
	name, cpf string,
	income float64,
	birthdate time.Time,
	children uint,
) (*Client, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if birthdate.After(time.Now()) {
		return nil, errors.New("birthdate cannot be in the future")
	}
	return &Client{
		Name:      name,
		Cpf:       cpf,
		Income:    income,
		BirthDate: birthdate,
		Children:  children,
	}, nil
}
