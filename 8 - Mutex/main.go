package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	counter := 0

	// goroutine 1
	go func() {
		mutex.Lock()
		counter++
		mutex.Unlock()
		fmt.Println("goroutine 1: counter =", counter)
	}()

	// goroutine 2
	go func() {
		mutex.Lock()
		counter++
		mutex.Unlock()
		fmt.Println("goroutine 2: counter =", counter)
	}()

	// wait for goroutines to finish
	time.Sleep(time.Second)
	fmt.Println("Final Counter", counter)
}
