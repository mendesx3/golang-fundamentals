package main

import "fmt"

type Calculator interface {
	calculateAverage(numbers []float64) float64
}

type CalculatorImpl struct {
}

func main() {
	numbers := []float64{10.5, 20.3, 15.7, 8.2, 12.6}
	avgCalculator := CalculatorImpl{}
	average := calculateUsingInterface(avgCalculator, numbers)
	fmt.Printf("Average: %.2f\n", average)

}

func (ac CalculatorImpl) calculateAverage(numbers []float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func calculateUsingInterface(calculator Calculator, numbers []float64) float64 {
	return calculator.calculateAverage(numbers)
}
