package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func main() {

	p := Person{"John", 25, "New York"}
	fmt.Println("Person:", p)
	fmt.Println("Person Detail:", p.Detail())

}

func (p Person) Detail() string {
	return fmt.Sprintf("%v (%v years old) from %v", p.Name, p.Age, p.City)
}
