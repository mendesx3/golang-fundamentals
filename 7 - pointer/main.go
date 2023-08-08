package main

import "fmt"

func main() {
	var a int = 42
	var pointer *int = &a

	fmt.Println("Value of a:", a)
	fmt.Println("Value of pointer:", pointer)

	*pointer = 100
	fmt.Println("Value of a:", a)
}
