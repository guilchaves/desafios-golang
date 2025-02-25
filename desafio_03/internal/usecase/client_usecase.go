package usecase

import (
	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
)

type ClientUseCase struct {
	repository entity.ClientRepositoryInterface
}

func NewClientUseCase(repository entity.ClientRepositoryInterface) *ClientUseCase {
	return &ClientUseCase{repository: repository}
}

func (uc *ClientUseCase) Create(client *entity.Client) error {
	return uc.repository.Save(client)
}

func (uc *ClientUseCase) GetClientByID(id int) (*entity.Client, error) {
	client, err := uc.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (uc *ClientUseCase) GetClients(page, limit int, sort string) ([]*entity.Client, error) {
	clients, err := uc.repository.FindAll(page, limit, sort)
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (uc *ClientUseCase) Update(client *entity.Client) error {
	return uc.repository.Update(client)
}

func (uc *ClientUseCase) Delete(id int) error {
	return uc.repository.Delete(id)
}
