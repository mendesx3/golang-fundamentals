package main

import "fmt"

func main() {

	aMap := make(map[string]int)

	aMap["one"] = 30
	aMap["two"] = 20
	aMap["three"] = 10

	exempl(aMap)

	if value, ok := aMap["two"]; ok {
		fmt.Println("value:", value)
		fmt.Println("exists:", ok)
	} else {
		fmt.Println("The key 'two' does not exist.")
	}
	delete(aMap, "two")
	fmt.Println("pós delete")
	fmt.Println("aMap:", aMap)
}

func exempl(aMap map[string]int) {
	fmt.Println("aMap:", aMap)
	fmt.Println("aMap 1", aMap["one"])
	fmt.Println("aMap 2", aMap["two"])
	fmt.Println("aMap 3", aMap["three"])
}
