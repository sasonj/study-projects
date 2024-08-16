package order

import "fmt"

type OrderItem struct {
	Name     string
	Quantity int
	Price    float64
}
type Order struct {
	ID    string
	Items []OrderItem
}

func AddOrder(orders map[string]Order, order Order) {

	orders[order.ID] = order
}

func CalculateTotal(order Order) float64 {
	var total float64 = 0

	for _, val := range order.Items {
		total += val.Price * float64(val.Quantity)
	}
	return total
}

func PrintOrders(orders map[string]Order) {
	for _, order := range orders {
		fmt.Println(order.ID, CalculateTotal(order))
	}
}
