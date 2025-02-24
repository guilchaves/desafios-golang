package database

import (

	"gorm.io/gorm"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
)

type ClientRepository struct {
	Db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{Db: db}
}

func (r *ClientRepository) Save(client *entity.Client) error {
	return r.Db.Create(client).Error

}


