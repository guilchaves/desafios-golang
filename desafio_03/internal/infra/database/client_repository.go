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

func (r *ClientRepository) FindByID(id int) (*entity.Client, error) {
	var client entity.Client
	err := r.Db.First(&client, "id = ?", id).Error

	return &client, err
}

func (r *ClientRepository) Update(client *entity.Client) error {
	_, err := r.FindByID(client.ID)
	if err != nil {
		return err
	}
	return r.Db.Save(client).Error
}

func (r *ClientRepository) Delete(id int) error {
	client, err := r.FindByID(id)
	if err != nil {
		return err
	}

	return r.Db.Delete(client).Error
}
