package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func main() {
	products := make(map[string]Product)

	products["p1"] = Product{"Pen", 1.99}
	products["p2"] = Product{"Pencil", 0.99}
	products["p3"] = Product{"Notebook", 5.99}

	if product, ok := products["p1"]; ok {
		fmt.Println("Product found:", product)
	} else {
		println("Product not found")
	}

	for key, product := range products {
		fmt.Printf("%v = %v\n", key, product)
	}
}
