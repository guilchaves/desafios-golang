package dto

import "time"

type CreateClientInputDTO struct {
	Name      string  `json:"name"`
	Cpf       string  `json:"cpf"`
	Income    float64 `json:"income"`
	BirthDate string  `json:"birthDate"`
	Children  uint    `json:"children"`
}

type UpdateClientInputDTO struct {
	Name      string  `json:"name"`
	Cpf       string  `json:"cpf"`
	Income    float64 `json:"income"`
	BirthDate string  `json:"birthDate"`
	Children  uint    `json:"children"`
}

type ClientOutputDTO struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Income    float64   `json:"income"`
	BirthDate time.Time `json:"birthDate"`
	Children  uint      `json:"children"`
}
