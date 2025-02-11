package service

import (
	"github.com/guilchaves/desafios-golang/desafio_01/internal/entity"
)

type OrderService struct {
	order           entity.Order
	shippingService *ShippingService
}

func NewOrderService(order entity.Order, shippingService *ShippingService) *OrderService {
	return &OrderService{order: order, shippingService: shippingService}
}

func (os *OrderService) CalculateTotal() float64 {
	shippingCost := os.shippingService.CalculateShipping()
	discount := os.order.BaseValue * (os.order.Discount / 100)
	discount = os.order.BaseValue - discount
	total := discount + shippingCost
	return total
}
