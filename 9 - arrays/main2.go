package main

import "fmt"

func main() {

	//slice
	numbers := []int{10, 20, 30, 40, 50}

	// print the slice
	fmt.Println("numbers:", numbers)

	// add values to slice
	numbers = append(numbers, 60)
	fmt.Println("numbers:", numbers)

}
