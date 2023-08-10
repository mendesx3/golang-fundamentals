package main

import "fmt"

func main() {

	add := func(x, y int) int {
		return x + y
	}

	result := add(1, 2)
	fmt.Println(result)

	double := func(x int) int {
		return x * 2
	}

	trible := func(x int) int {
		return x * 3
	}

	processNumbers(double)
	processNumbers(trible)
}

func processNumbers(fn func(x int) int) {
	numbers := []int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		fmt.Println("processing number", fn(number))
	}
}
