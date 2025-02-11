package service

import (
    "testing"
    "github.com/guilchaves/desafios-golang/desafio_01/internal/entity"
    "github.com/stretchr/testify/assert"
)


func TestCalculateOrder(t *testing.T) {
    order := entity.Order{
        Code:      1034,
        BaseValue: 150.0,
        Discount:  20.0,
    }
    shippingService := NewShippingService(order)
    os := NewOrderService(order, shippingService)
    orderTotal := os.CalculateTotal()

    assert.Equal(t, 132.0, orderTotal)
    order = entity.Order{
        Code:      2282,
        BaseValue: 800.0,
        Discount:  10.0,
    }

    shippingService = NewShippingService(order)
    os = NewOrderService(order, shippingService)
    orderTotal = os.CalculateTotal()

    assert.Equal(t, 720.0, orderTotal)

}
