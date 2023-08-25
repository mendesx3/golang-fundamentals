package main

import "fmt"

func CalculateSquare(number int, resultChan chan int) {
	result := number * number
	resultChan <- result
}

func main() {

	resultChannel := make(chan int)

	numGoRoutines := 5

	for i := 0; i < numGoRoutines; i++ {
		go CalculateSquare(i, resultChannel)
	}

	for i := 0; i < numGoRoutines; i++ {
		result := <-resultChannel
		fmt.Printf("Result: %d\n", result)
	}

	close(resultChannel)

}
