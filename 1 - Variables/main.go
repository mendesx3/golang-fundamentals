package main

import "fmt"

func main() {
	// Calling the functions you've defined to demonstrate the examples.
	example()
	example2()
	example3()
	example4()
	example5()
	example6()
	example7()
}

// example demonstrates a standard variable declaration and assignment.
func example() {
	var name string   // Variable declaration of type string
	name = "John Doe" // Value assignment to the variable
	fmt.Println("Hello, " + name)
}

// example2 showcases declaring a variable with initialization.
func example2() {
	var age int = 30 // Variable declaration and initialization on the same line
	fmt.Println("Age: ", age)
}

// example3 employs the shorthand variable declaration.
func example3() {
	name := "John Doe" // Shorthand declaration and initialization
	fmt.Println("Hello, " + name)
}

// example4 illustrates type conversion from int to float64.
func example4() {
	var x int = 10             // Declaration of an int variable
	var y float64 = float64(x) // Type conversion from int to float64
	fmt.Println("y: ", y)
}

// example5 presents a way to declare multiple variables.
func example5() {
	// Multiple variables declaration with their values
	var (
		a = 1
		b = 2
		c = 3
	)
	fmt.Println("a: ", a, "b: ", b, "c: ", c)
}

// example6 defines and uses a constant.
func example6() {
	const PI = 3.14 // Constant declaration
	fmt.Println("PI: ", PI)
}

// GlobalVar Definition of global and package level variables.
var GlobalVar = "global!" // Global scope variable
var pkgVar = "pkg"        // Package scope variable

// example7 demonstrates the use of local, package, and global variables.
func example7() {
	var localVar = "local!" // Local scope variable
	fmt.Println(localVar, pkgVar, GlobalVar)
}
