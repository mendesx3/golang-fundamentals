package main

import "fmt"

type Item interface {
	Price() float64
}

type PhysicalProduct struct {
	Name      string
	PriceItem float64
	Wight     float64
}

type DigitalProduct struct {
	Name      string
	PriceItem float64
}

func (p Service) Price() float64 {
	return p.PriceItem
}

type Service struct {
	Name      string
	PriceItem float64
}

func CalculateOrderTotal(Items []Item) float64 {
	total := 0.0
	for _, item := range Items {
		total += item.Price()
	}
	return total
}

func (d DigitalProduct) Price() float64 {
	return d.PriceItem
}

func (p PhysicalProduct) Price() float64 {
	return p.PriceItem
}

func main() {

	items := []Item{
		PhysicalProduct{
			Name:      "Book Go",
			PriceItem: 100.00,
			Wight:     5.5,
		},
		DigitalProduct{
			Name:      "Book Java",
			PriceItem: 9.99,
		},
		Service{
			Name:      "Book Python",
			PriceItem: 0.5,
		},
	}
	total := CalculateOrderTotal(items)
	fmt.Printf("Total order amount: %.2f", total)

}
