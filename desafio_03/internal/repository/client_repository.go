package repository

import (
	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db}
}

func (r *ClientRepository) Create(client *entity.Client) error {
	return r.db.Create(client).Error
}

func (r *ClientRepository) FindByID(id uint) (*entity.Client, error) {
	var client entity.Client
	err := r.db.First(&client, id).Error
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *ClientRepository) FindAll(limit, offset int) ([]entity.Client, error) {
	var clients []entity.Client
	err := r.db.Limit(limit).Offset(offset).Find(&clients).Error
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (r *ClientRepository) Update(client *entity.Client) error {
	return r.db.Model(&entity.Client{}).Where("id = ?", client.ID).Updates(client).Error
}

func (r *ClientRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Client{}, id).Error
}
