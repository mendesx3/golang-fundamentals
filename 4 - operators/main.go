package main

import "fmt"

func main() {
	exemple()
	exemple2()
	exemple3()
	exemple4()

}

func exemple() {
	a, b := 10, 3
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Remainder:", a%b)
}

func exemple2() {
	a, b := 10, 3
	fmt.Println("AND:", a&b)
	fmt.Println("OR:", a|b)
	fmt.Println("XOR:", a^b)
	fmt.Println("AND NOT:", a&^b)
	fmt.Println("Left Shift:", a<<2)
	fmt.Println("Right Shift:", a>>2)
}

func exemple3() {
	a, b := 10, 3
	fmt.Println("Equal to:", a == b)
	fmt.Println("Not equal to:", a != b)
	fmt.Println("Less than:", a < b)
	fmt.Println("Less than or equal to:", a <= b)
	fmt.Println("Greater than:", a > b)
	fmt.Println("Greater than or equal to:", a >= b)
}

func exemple4() {
	a := 10
	a = 5  // a = 5
	a += 3 // a = a + 3 = 8
	a -= 2 // a = a - 2 = 6
	a *= 4 // a = a * 4 = 24
	a /= 3 // a = a / 3 = 8
	a %= 5 // a = a % 5 = 3

	fmt.Println("Final value of a:", a) // Output: 3
}
