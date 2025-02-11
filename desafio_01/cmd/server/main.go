package main

import (
	"fmt"
	"os"

	"github.com/guilchaves/desafios-golang/desafio_01/internal/entity"
	"github.com/guilchaves/desafios-golang/desafio_01/internal/service"
)

func main() {
	var code int
	var baseValue, discount float64

	fmt.Print("Enter the order code: ")
	_, err := fmt.Scanf("%d", &code)

	if err != nil {
		fmt.Println("Error reading order code:", err)
		os.Exit(1)
	}

	fmt.Print("Enter the base order value: ")
	_, err = fmt.Scanf("%f", &baseValue)
	if err != nil {
		fmt.Println("Error reading base order value:", err)
		os.Exit(1)
	}

	fmt.Print("Enter the discount percentage: ")
	_, err = fmt.Scanf("%f", &discount)
	if err != nil {
		fmt.Println("Error reading discount percentage:", err)
		os.Exit(1)
	}

	order, err := entity.NewOrder(code, baseValue, discount)
	if err != nil {
		panic(err)
	}

	shippingService := service.NewShippingService(*order)
	os := service.NewOrderService(*order, shippingService)
	orderTotal := os.CalculateTotal()
	fmt.Printf("Order code: %d\n", order.Code)
	fmt.Printf("Order total: R$ %.2f\n", orderTotal)
}
