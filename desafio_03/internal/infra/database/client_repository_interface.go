package database 

import "github.com/guilchaves/desafios-golang/desafio_03/internal/entity"

type ClientRepository interface {
	Create(client *entity.Client) error
	FindByID(id uint) (*entity.Client, error)
	FindAll(limit, offset int) ([]entity.Client, error)
	Update(client *entity.Client) error
	Delete(id uint) error
}

