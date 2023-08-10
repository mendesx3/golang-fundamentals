package main

import "fmt"

func main() {

	age := 18

	if age <= 18 {
		fmt.Println("You are an adult")
	} else if age < 60 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are an adult")
	}

	for i := 0; i < 10; i++ {
		fmt.Println("counting...", i)

		day := "Monday"
		switch day {
		case "Monday":
			fmt.Println("Today is Monday")
		case "Tuesday":
			fmt.Println("Today is Tuesday")
		case "Wednesday":
			fmt.Println("Today is Wednesday")
		case "Thursday":
			fmt.Println("Today is Thursday")
		case "Friday":
			fmt.Println("Today is Friday")
		case "Saturday":
			fmt.Println("Today is Saturday")
		case "Sunday":
			fmt.Println("Today is Sunday")
		default:
			fmt.Println("Today is a new day")
		}
	}
}
