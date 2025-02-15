package entity

import (
	"time"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/validator"
)

type Client struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CPF       string
	Income    float64
	BirthDate time.Time
	Children  int
}

func NewClient(
	name, cpf string,
	income float64,
	birthDate time.Time,
	children int,
) (*Client, error) {
	if err := validator.ValidateClientName(name); err != nil {
		return nil, err
	}

	if err := validator.ValidateCPF(cpf); err != nil {
		return nil, err
	}

	if err := validator.ValidateIncome(income); err != nil {
		return nil, err
	}

	if err := validator.ValidateClientBirthDate(birthDate); err != nil {
		return nil, err
	}

	if err := validator.ValidateChildren(children); err != nil {
		return nil, err
	}

	return &Client{
		Name:      name,
		CPF:       cpf,
		Income:    income,
		BirthDate: birthDate,
		Children:  children,
	}, nil
}
