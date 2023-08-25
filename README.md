# Go (Golang) Fundamentals

In this document, we will list the fundamentals of the Go programming language (Golang).

## Variables
- In Go, variables are declared with the `var` keyword.
- You can specify the type of a variable or let the compiler automatically infer it.
- Local variables are declared using `:=`.

## Data Types
- Go has basic data types such as integers, floats, strings, and booleans.
- It also offers composite types like arrays, slices, maps, structs, and interfaces.

## Functions
- Functions are defined using the `func` keyword.
- Parameters and return types are specified in the function signature.

## Control Structures
- Go supports control structures like `if`, `else`, `for`, `switch`, and `select`.

## Packages
- Go organizes code into packages, and a Go program is composed of one or more packages.

## Goroutines
- Goroutines are lightweight units of execution that enable concurrency in Go.
- They are created using `go`.

## Channels
- Channels are used for safe communication between goroutines.
- They are created using `make(chan type)`.

## Data Structures
- Go provides data types such as arrays, slices, maps, and structs for structuring data.

## Pointers
- Pointers are used to store the memory address of a variable.
- They are declared with the `&` operator and accessed with the `*` operator.

## Interfaces
- Interfaces allow defining contracts that data types must adhere to.
- Interfaces are implemented implicitly in Go.

## Defer and Panic
- `defer` is used to delay the execution of a function until the containing function returns.
- `panic` is used to create runtime errors.