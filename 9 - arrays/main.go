package main

import "fmt"

func main() {

	var numbers [5]int

	// add values to the array
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	// print the array
	fmt.Println("numbers:", numbers)
	fmt.Println("numbers[0]:", numbers[0])
	fmt.Println("numbers[4]:", numbers[4])
}
