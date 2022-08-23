package models

import "fmt"

type Order struct {
	Apples  int
	Oranges int
	Cost    float64
}

type OrderService struct{}

func (orderService OrderService) MakeOrder(order *Order) Order {
	order.Cost = (float64(order.Apples/2 + order.Apples%2) * 0.6) + (float64(order.Oranges/3) * 0.5) + (float64(order.Oranges%3) * 0.25)
	fmt.Println(order.Cost)
	return *order
}
