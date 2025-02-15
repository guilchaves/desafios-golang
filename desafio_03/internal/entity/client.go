package entity

import (
	"time"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/validator"
)

type Client struct {
	ID        uint      `gorm:"primaryKey; autoIncrement"`
	Name      string    `gorm:"column:name;not null"`
	CPF       string    `gorm:"column:cpf;unique;not null"`
	Income    float64   `gorm:"column:income;not null"`
	BirthDate time.Time `gorm:"column:birth_date;not null"`
	Children  int       `gorm:"column:children;not null"`
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
