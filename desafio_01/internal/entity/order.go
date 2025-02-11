package entity

import "errors"

type Order struct {
	Code      int
	BaseValue float64
	Discount  float64
}

var (
	ErrCodeIsRequired      = errors.New("code is required")
	ErrBaseValueIsRequired = errors.New("base value is required")
	ErrBaseValueIsInvalid  = errors.New("base value is invalid")
	ErrDiscountIsRequired  = errors.New("discount is required")
	ErrDiscountIsIsInvalid = errors.New("discount is invalid")
)

func NewOrder(code int, baseValue, discount float64) (*Order, error) {
	order := &Order{
		Code:      code,
		BaseValue: baseValue,
		Discount:  discount,
	}

	err := order.Validate()
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) Validate() error {
	if o.Code == 0 {
		return ErrCodeIsRequired
	}
	if o.BaseValue == 0 {
		return ErrBaseValueIsRequired
	}
	if o.BaseValue < 0 {
		return ErrBaseValueIsInvalid
	}
	if o.Discount == 0 {
		return ErrDiscountIsRequired
	}
	if o.Discount < 0 {
		return ErrDiscountIsIsInvalid
	}
	return nil
}
