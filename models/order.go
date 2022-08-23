package models

import (
	"database/sql"
	"fmt"
)

type Order struct {
	ID      int
	Apples  int
	Oranges int
	Cost    float64
}

type OrderService struct {
	DB *sql.DB
}

func (orderService OrderService) MakeOrder(order *Order) error {
	order.Cost = (float64(order.Apples/2+order.Apples%2) * 0.6) + (float64(order.Oranges/3) * 0.5) + (float64(order.Oranges%3) * 0.25)
	fmt.Println(order.Cost)

	row := orderService.DB.QueryRow(`
		insert into Orders(apples, oranges, cost)
		values($1, $2, $3)
		returning id
	`, order.Apples, order.Oranges, order.Cost)

	err := row.Scan(&order.ID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("could not add order to database")
	}

	return nil
}

func (orderService OrderService) GetById(id int) (*Order, error) {

	order := Order{}

	row := orderService.DB.QueryRow(`
		select id, apples, oranges, cost from Orders
		where id=$1
	`, id)
	err := row.Scan(&order.ID, &order.Apples, &order.Oranges, &order.Cost)
	if (err == sql.ErrNoRows) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("there was an issue finding this order")
	}
	return &order, nil
}

func (orderService OrderService) GetAll() (*[]*Order, error) {

	orders := make([]*Order, 0)

	rows, err := orderService.DB.Query(`
	select * from Orders
	`)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("could not reach orders")
	}
	for rows.Next() {
		order := new(Order)
		err := rows.Scan(&order.ID, &order.Apples, &order.Oranges, &order.Cost)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	return &orders, nil
}
