package usecase

import (
	"errors"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/repository"

	"time"
)

var (
	ErrBirthDateInvalid = errors.New("data de nascimento inválida")
)

type ClientUsecase struct {
	clientRepo *repository.ClientRepository
}

func NewClientUsecase(clientRepo *repository.ClientRepository) *ClientUsecase {
	return &ClientUsecase{
		clientRepo: clientRepo,
	}
}

func (uc *ClientUsecase) Create(
	name, cpf string,
	income float64,
	birthDateStr string,
	children int,
) (*entity.Client, error) {
	birthDate, err := time.Parse("2006-01-02", birthDateStr)
	if err != nil {
		return nil, ErrBirthDateInvalid
	}

	client, err := entity.NewClient(name, cpf, income, birthDate, children)
	if err != nil {
		return nil, err
	}

	err = uc.clientRepo.Create(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}
