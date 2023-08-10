package main

import "fmt"

type (
	Item struct {
		Name     string
		Price    float64
		Quantity int
	}

	Order struct {
		items    []Item
		Discount float64
	}
)

func (o Order) Total() float64 {
	total := 0.0
	for _, item := range o.items {
		total += item.Price * float64(item.Quantity)
	}
	if o.Discount > 0 {
		total -= total * o.Discount
	}

	return total
}

func main() {

	order := Order{
		items: []Item{
			{Name: "Pen", Price: 5.20, Quantity: 2},
			{Name: "Pencil", Price: 4.80, Quantity: 1},
		},
		Discount: 0.10,
	}
	total := order.Total()
	fmt.Printf("Total: $%.2f\n", total)
}
