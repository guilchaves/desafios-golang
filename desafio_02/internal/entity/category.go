package entity

import "errors"

type CategoryDescription string

const (
	Oficina CategoryDescription = "Oficina"
	Curso   CategoryDescription = "Curso"
)

var (
	ErrCategoryIDIsRequired          = errors.New("id is required")
	ErrCategoryDescriptionIsRequired = errors.New("description is required")
	ErrCategoryDescriptionIsInvalid  = errors.New("description is invalid")
	ErrCategoryActivitiesIsRequired  = errors.New("activities is required")
)

type Category struct {
	ID          int                 `gorm:"primaryKey"`
	Description CategoryDescription `json:"description"`
	Activities  []Activity          `gorm:"foreignKey:CategoryID"`
}

func NewCategory(id int, description CategoryDescription) (*Category, error) {
	category := &Category{
		ID:          id,
		Description: description,
		Activities:  []Activity{},
	}
	if err := category.Validate(); err != nil {
		return nil, err
	}
	return category, nil
}

func (c *Category) Validate() error {
	if c.ID == 0 {
		return ErrCategoryIDIsRequired
	}
	if c.Description == "" {
		return ErrCategoryDescriptionIsRequired
	}
	if c.Description != Oficina && c.Description != Curso {
		return ErrCategoryDescriptionIsInvalid
	}

	return nil
}
