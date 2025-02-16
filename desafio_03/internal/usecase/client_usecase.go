package usecase

import (
	"errors"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/infra/database"

	"time"
)

var (
	ErrBirthDateInvalid = errors.New("data de nascimento inválida")
)

type ClientUsecase struct {
	clientRepo database.ClientRepository
}

func NewClientUsecase(clientRepo database.ClientRepository) *ClientUsecase {
	return &ClientUsecase{
		clientRepo: clientRepo,
	}
}

func (uc *ClientUsecase) CreateClient(
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

func (uc *ClientUsecase) GetClients() ([]entity.Client, error) {
	return uc.clientRepo.FindAll(10, 0)
}

func (uc *ClientUsecase) GetClientByID(id uint) (*entity.Client, error) {
	return uc.clientRepo.FindByID(id)
}

func (uc *ClientUsecase) UpdateClient(client *entity.Client) error {
	return uc.clientRepo.Update(client)
}

func (uc *ClientUsecase) DeleteClient(id uint) error {
	return uc.clientRepo.Delete(id)
}
