package main

import "fmt"

type (
	Person struct {
		Name  string
		Email string
		City  string
	}

	Employee struct {
		Person
		EmployeeID int
		Department string
		Salary     float64
	}
)

func main() {
	employee := Employee{
		Person: Person{
			Name:  "John Doe",
			Email: "@gmail.com",
			City:  "New York",
		},
		EmployeeID: 123,
		Department: "IT",
		Salary:     1000.00,
	}
	fmt.Println("Employee Name: " + employee.Name)
	fmt.Println("Employee Email: " + employee.Email)
	fmt.Println("Employee City: " + employee.City)
	fmt.Println("Employee ID: ", employee.EmployeeID)
	fmt.Println("Employee Department: " + employee.Department)
	fmt.Println("Employee Salary: ", employee.Salary)
}
