package main

import "fmt"

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	c := Circle{
		Radius: 5,
	}
	area := c.Area()
	fmt.Printf("Area of the circle: %.2f\n", area)
}
