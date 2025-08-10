package basics

import "fmt"

var middleName string = "Middle"

func main() {
	var age int
	var name string = "John Doe"
	var name1 = "Jane Doe"

	count := 10
	lastName := "Smith"

	// Default values
	// Numeric types: 0
	// Boolean types: false
	// string Type: ""
	// Pointers, slices, maps, functions and structs: nil

	// scope
	fmt.Println(middleName)
}

func printname() {
	firstName := "Minstrel"
	fmt.Println(firstName)
	fmt.Println(middleName)
}
