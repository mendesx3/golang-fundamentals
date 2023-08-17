package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64) bool
}

type CreditCard struct {
	Number       string
	Expiration   string
	CVV          string
	AccountOwner string
}

type Paypal struct {
	Username string
	Password string
}

func (p Paypal) Pay(amount float64) bool {
	fmt.Printf("Processing paypal payment of %.2f\n", amount)
	return true
}

func (c CreditCard) Pay(amount float64) bool {
	fmt.Printf("Processing credit card payment of %.2f\n", amount)
	return true
}

func processOayment(method PaymentMethod, amount float64) {
	if method.Pay(amount) {
		fmt.Println("Payment successfull")
	} else {
		fmt.Println("Payment failed")
	}
}

func main() {

	creditCard := CreditCard{
		Number:       "1234 1234 1234 1234",
		Expiration:   "12/21",
		CVV:          "123",
		AccountOwner: "John Doe",
	}

	paypal := Paypal{
		Username: "@gmail.com",
		Password: "123456",
	}
	amount := 500.00

	processOayment(creditCard, amount)
	processOayment(paypal, amount)

}
