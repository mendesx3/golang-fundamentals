package main

import "fmt"

type (
	Employee struct {
		Name       string
		Position   string
		Salary     float64
		Department string
	}
	Company struct {
		Name         string
		Registration string
		Employees    []Employee
	}
)

func (e *Employee) IncreaseSalary(p float64) {
	e.Salary = e.Salary + (e.Salary * p / 100)
}

func (c *Company) HireEmployee(e Employee) {
	c.Employees = append(c.Employees, e)
}

func (c *Company) FireEmployee(e Employee) {
	for i, v := range c.Employees {
		if v.Name == e.Name {
			c.Employees = append(c.Employees[:i], c.Employees[i+1:]...)
			break
		}
	}
}

func (c Company) GetEmployees() {
	fmt.Println("Employees of Company", c.Name)
	for _, e := range c.Employees {
		fmt.Println("Name:", e.Name)
		fmt.Println("Position:", e.Position)
		fmt.Println("Salary:", e.Salary)
		fmt.Println("Department:", e.Department)
	}
}

func main() {
	// create company
	c := Company{
		Name:         "ABC",
		Registration: "123456789",
	}

	// create employee
	e1 := Employee{
		Name:       "John",
		Position:   "Manager",
		Salary:     1000,
		Department: "IT",
	}
	e2 := Employee{
		Name:       "Jane",
		Position:   "Staff",
		Salary:     500,
		Department: "IT",
	}

	fmt.Println("hiring employees...")
	c.HireEmployee(e1)
	c.HireEmployee(e2)

	c.GetEmployees()
	fmt.Println("----------------------------")

	fmt.Println("increasing salary...")
	e1.IncreaseSalary(10)
	e2.IncreaseSalary(50)
	c.GetEmployees()

	fmt.Println("firing employee...")
	c.FireEmployee(e1)
	c.GetEmployees()

}
