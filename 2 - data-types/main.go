package main

import "fmt"

func main() {
	// Calling each of the demonstration functions.
	exempleInt()
	exempleFloat()
	exempleString()
	exempleBool()
	exempleArray()
}

// exempleInt showcases the use of integer data types in Go.
func exempleInt() {
	var a int = 10   // Declaring an integer of type 'int'
	var b int64 = 20 // Declaring a 64-bit integer
	fmt.Println(a)
	fmt.Println(b)
}

// exempleFloat demonstrates the use of floating-point numbers in Go.
func exempleFloat() {
	var a float32 = 10.5 // Declaring a 32-bit floating-point number
	var b float64 = 20.5 // Declaring a 64-bit floating-point number
	fmt.Println(a)
	fmt.Println(b)
}

// exempleString highlights the use of string data types in Go.
func exempleString() {
	var a string = "Hello" // Declaring a string using explicit type annotation
	var b = "Hello"        // Declaring a string without explicit type (type inferred from value)
	fmt.Println(a)
	fmt.Println(b)
}

// exempleBool demonstrates the use of boolean data types in Go.
func exempleBool() {
	var a bool = true // Declaring a boolean with the value 'true'
	var b = false     // Declaring a boolean without explicit type (type inferred from value)
	fmt.Println(a)
	fmt.Println(b)
}

// exempleArray illustrates the creation and manipulation of arrays in Go.
func exempleArray() {
	var a [5]int // Declaring an array that can hold 5 integers, initialized to zero values by default
	a[2] = 7     // Assigning the value 7 to the third position (0-indexed) of the array
	fmt.Println(a)
}
