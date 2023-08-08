package main

import (
	"fmt"
	"sync"
)

type Account struct {
	ID      int
	Balance float64
	Mutex   sync.Mutex
}

func Transfer(from, to *Account, amount float64) {
	from.Mutex.Lock()
	defer from.Mutex.Unlock()

	to.Mutex.Lock()
	defer to.Mutex.Unlock()

	if from.Balance >= amount {
		from.Balance -= amount
		to.Balance += amount
		fmt.Println("Transfered", amount, "from", from.ID, "to", to.ID)
	} else {
		fmt.Println("Insufficient balance in", from.ID)
	}

}

func main() {
	// create two accounts
	account1 := Account{ID: 1, Balance: 1000}
	account2 := Account{ID: 2, Balance: 3000}

	// perform transfer from account1 to account2
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		Transfer(&account1, &account2, 500)
	}()
	go func() {
		defer wg.Done()
		Transfer(&account2, &account1, 200)
	}()
	wg.Wait()
	fmt.Println("Final balance of account1:", account1.Balance)
	fmt.Println("Final balance of account2:", account2.Balance)

}
