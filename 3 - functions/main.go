package main

import (
	"fmt"
)

func main() {
	example("John", "20")
	r := example2(5, 2)
	fmt.Println(r)

	r1, r2 := example3(5, 2)
	fmt.Println(r1, r2)

	r3, r4 := example4(5, 2)
	fmt.Println(r3, r4)

	r5, r6 := example5(5, 2)
	fmt.Println(r5, r6)

	example6(1, 2, 3, 4, 5)

	result := func(a, b int) int {
		return a + b
	}
	fmt.Println(result(1, 2))

	// Additional Example 1: Finding the longest and shortest strings in a list
	list := []string{"apple", "banana", "orange", "grape"}
	longest, shortest := findLongestShortest(list)
	fmt.Println("Longest string:", longest)
	fmt.Println("Shortest string:", shortest)

	// Additional Example 2: Calculate the year of birth based on age
	person := Person{Name: "Alice", Age: 30}
	yearOfBirth := person.calculateYearOfBirth()
	fmt.Println(person.Name, "was born in the year:", yearOfBirth)

	// Additional Example 3: Recursive function to calculate the nth Fibonacci number
	n := 8
	fmt.Println("The", n, "th Fibonacci number is:", fibonacci(n))

	// Additional Example 4: Interface and implementations for different shapes
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 6, Height: 4}
	triangle := Triangle{Base: 5, Height: 3}
	shapes := []Shape{circle, rectangle, triangle}
	for _, shape := range shapes {
		fmt.Println("Area of", shape.Name(), ":", shape.Area())
	}
}

func example(name, age string) {
	fmt.Println("Hello, " + name)
	fmt.Println("You are " + age + " years old")
}

func example2(a, b int) int {
	return a - b
}

func example3(p1, p2 int) (int, int) {
	return p1, p2
}

func example4(p1, p2 int) (r1, r2 int) {
	r1 = p1 + p2
	r2 = p1 - p2
	return
}

func example5(p1, p2 int) (r1, r2 int) {
	r1 = p1 + p2
	r2 = p1 - p2
	return r1, r2
}

func example6(numbers ...int) {
	count := 0
	for _, number := range numbers {
		count += number
	}
	fmt.Println(count)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) calculateYearOfBirth() int {
	currentYear := 2023 // Assuming the current year is 2023
	return currentYear - p.Age
}

func findLongestShortest(list []string) (longest, shortest string) {
	longest = ""
	shortest = list[0]
	for _, str := range list {
		if len(str) > len(longest) {
			longest = str
		}
		if len(str) < len(shortest) {
			shortest = str
		}
	}
	return
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Additional Example 4: Interface and implementations for different shapes
type Shape interface {
	Name() string
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Name() string {
	return "Circle"
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Name() string {
	return "Rectangle"
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Name() string {
	return "Triangle"
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
