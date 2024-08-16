package main

import (
	"fmt"
	"kkkafesoft/order"
)

func main() {
	fmt.Println("kkkafesoft")

	order1 := order.Order{
		ID: "first",
		Items: []order.OrderItem{
			{
				Name:     "apple",
				Quantity: 5,
				Price:    100,
			},
			{
				Name:     "grape",
				Quantity: 10,
				Price:    200,
			},
		},
	}

	order2 := order.Order{
		ID: "Second",
		Items: []order.OrderItem{
			{
				Name:     "apple",
				Quantity: 5,
				Price:    200,
			},
			{
				Name:     "grape",
				Quantity: 10,
				Price:    400,
			},
		},
	}

	OrderMap := make(map[string]order.Order)

	order.AddOrder(OrderMap, order1)
	order.AddOrder(OrderMap, order2)

	order.PrintOrders(OrderMap)

}
