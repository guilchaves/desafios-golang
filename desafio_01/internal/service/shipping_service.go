package service

import "github.com/guilchaves/desafios-golang/desafio_01/internal/entity"

type ShippingService struct {
	order entity.Order
}

func NewShippingService(order entity.Order) *ShippingService {
	return &ShippingService{order: order}
}

func (ss *ShippingService) CalculateShipping() float64 {
	if ss.order.BaseValue < 100 {
		return 20.0
	}
	if ss.order.BaseValue >= 100.0 && ss.order.BaseValue < 200.0 {
		return 12.00
	}
	return 00.0
}
