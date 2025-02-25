package dto

import "time"

type ClientInputDTO struct {
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Income    float64   `json:"income"`
	BirthDate time.Time `json:"birthDate"`
	Children  uint      `json:"children"`
}

type ClientOutputDTO struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Income    float64   `json:"income"`
	BirthDate time.Time `json:"birthDate"`
	Children  uint      `json:"children"`
}
