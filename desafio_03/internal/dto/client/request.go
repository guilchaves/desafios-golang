package client

type CreateClientRequest struct {
	Name      string  `json:"name" validate:"required"`
	CPF       string  `json:"cpf" validate:"required,cpf"`
	Income    float64 `json:"income"`
	BirthDate string  `json:"birth_date" validate:"required,datetime=2006-01-02"`
	Children  int     `json:"children"`
}

type UpdateClientRequest struct {
	Name      string  `json:"name" validate:"omitempty,min=1"`
	Income    float64 `json:"income" validate:"omitempty, min=0"`
	BirthDate string  `json:"birth_date" validate:"omitempty,datetime=2006-01-02"`
	Children  int     `json:"children" validate:"omitempty,min=0"`
}
