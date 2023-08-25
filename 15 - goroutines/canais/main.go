package main

import "fmt"

func main() {

	intChannel := make(chan int)

	go func() {
		intChannel <- 42
	}()

	values := <-intChannel
	fmt.Println(values)
}
