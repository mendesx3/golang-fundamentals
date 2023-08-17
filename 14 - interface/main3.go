package main

import "fmt"

type Wallet interface {
	Balance() float64
	Send(amount float64, recipient string) error
}

type EthereumWallet struct {
	Address       string
	BalanceWallet float64
}

func (e *EthereumWallet) Balance() float64 {
	return e.BalanceWallet
}

func (e *EthereumWallet) Send(amount float64, recipient string) error {
	if amount <= e.BalanceWallet {
		e.BalanceWallet -= amount
		fmt.Printf("Sent %.2f ETH to %s\n", amount, recipient)
		return nil
	}
	return fmt.Errorf("insufficient balance")
}

type BitcoinWallet struct {
	Address       string
	BalanceWallet float64
}

func (b *BitcoinWallet) Balance() float64 {
	return b.BalanceWallet
}

func (b *BitcoinWallet) Send(amount float64, recipient string) error {
	if amount <= b.BalanceWallet {
		b.BalanceWallet -= amount
		fmt.Printf("Sent %.2f BTC to %s\n", amount, recipient)
		return nil
	}
	return fmt.Errorf("insufficient balance")
}

func ProcessWallet(w Wallet, amount float64, recipient string) {
	fmt.Printf("Wallet Address: %s\n", w.(Wallet).Balance())
	fmt.Printf("Balance: %.2f\n", w.Balance())

	err := w.Send(amount, recipient)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func main() {
	ethWallet := &EthereumWallet{
		Address:       "0x3232",
		BalanceWallet: 100,
	}

	bitWallet := &BitcoinWallet{
		Address:       "0x3232",
		BalanceWallet: 100,
	}

	ProcessWallet(ethWallet, 10, "0x1234")
	ProcessWallet(bitWallet, 10, "0x1234")

}
