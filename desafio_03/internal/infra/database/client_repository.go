package database

import (
	"strings"

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
	if err != nil {
		return nil, err
	}
	return &client, err
}

func (r *ClientRepository) FindAll(page, limit int, sort string) ([]*entity.Client, error) {
	var clients []*entity.Client

	sort = strings.ToLower(sort)
	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	query := r.Db.Order("id " + sort)
	if page > 0 && limit > 0 {
		query = query.Limit(limit).Offset((page - 1) * limit)
	}

	err := query.Find(&clients).Error
	return clients, err
}

func (r *ClientRepository) Update(client *entity.Client) error {
	result := r.Db.Model(&entity.Client{}).Where("id = ?", client.ID).Updates(client)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error

}

func (r *ClientRepository) Delete(id int) error {
	result := r.Db.Delete(&entity.Client{}, id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
