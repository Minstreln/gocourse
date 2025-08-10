package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// e.g CalculateArea, UserInfo
	// Structs, Interfaces, Enums

	// snake_case
	// Eg user_id, first_name
	// Variables, Functions, Package names

	// UPPERCASE
	// used exclusively for constants

	// mixedCase
	// e.g calculateArea, userInfo

	const MAXRETRIES = 5
	var employeeID = 1001
	fmt.Println("Employee ID", employeeID)
}