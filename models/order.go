package models

import "fmt"

type Order struct {
	Apples  int
	Oranges int
	Cost    float64
}

type OrderService struct{}

func (orderService OrderService) MakeOrder(order *Order) Order {
	order.Cost = (float64(order.Apples) * 0.6) + (float64(order.Oranges) * 0.25)
	fmt.Println(order.Cost)
	return *order
}
