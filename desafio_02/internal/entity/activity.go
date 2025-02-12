package entity

import "errors"

type Activity struct {
	ID          int         `gorm:"primaryKey"`
	Name        string      `gorm:"not null"`
	Description string      `gorm:"not null"`
	Price       float64     `gorm:"not null"`
	CategoryID  int         `gorm:"not null"                                       json:"category_id"`
	Category    Category    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TimeBlocks  []TimeBlock `gorm:"foreignKey:ActivityID"`
}

var (
	ErrActivityNameIsRequired        = errors.New("name is required")
	ErrActivityDescriptionIsRequired = errors.New("description is required")
	ErrActivityPriceIsRequired       = errors.New("price is required")
	ErrActivityPriceIsInvalid        = errors.New("price is invalid")
	ErrActivityCategoryIsRequired    = errors.New("category id is required")
)

func NewActivity(
	name, description string,
	price float64,
	categoryID int,
	timeBlocks []TimeBlock,
) (*Activity, error) {
	activity := &Activity{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		TimeBlocks:  timeBlocks,
	}
	if err := activity.Validate(); err != nil {
		return nil, err
	}
	return activity, nil
}

func (a *Activity) Validate() error {
	if a.Name == "" {
		return ErrActivityNameIsRequired
	}
	if a.Description == "" {
		return ErrActivityDescriptionIsRequired
	}
	if a.Price == 0 {
		return ErrActivityPriceIsRequired
	}
	if a.Price < 0 {
		return ErrActivityPriceIsInvalid
	}
	if a.CategoryID == 0 {
		return ErrActivityCategoryIsRequired
	}

	for _, block := range a.TimeBlocks {
		if err := block.Validate(); err != nil {
			return err
		}
	}

	return nil
}
