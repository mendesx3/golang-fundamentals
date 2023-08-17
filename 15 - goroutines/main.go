package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i < 10; i++ {
		println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	go printNumbers()

	for i := 1; i <= 5; i++ {
		fmt.Printf("main-%d ", i)
		time.Sleep(200 * time.Millisecond) // Simulando uma operação demorada
	}

}
