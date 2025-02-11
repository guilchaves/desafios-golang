package service

import (
    "testing"
    "github.com/guilchaves/desafios-golang/desafio_01/internal/entity"
    "github.com/stretchr/testify/assert"
)

func TestCalculateShipping(t *testing.T) {
    order := entity.Order{
        Code:      1034,
        BaseValue: 150.0,
        Discount:  20.0,
    }
    ss := NewShippingService(order)
    shippingCost := ss.CalculateShipping()
    assert.Equal(t, 12.0, shippingCost)

    order = entity.Order{
        Code:      2282,
        BaseValue: 800.0,
        Discount:  10.0,
    }
    ss = NewShippingService(order)
    shippingCost = ss.CalculateShipping()
    assert.Equal(t, 0.0, shippingCost)
}

