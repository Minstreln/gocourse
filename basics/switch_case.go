package basics

import "fmt"


func main() {

	// switch statements in Go (switch case, default, fallthrough)
	fruit := "apple"

	switch fruit {
	case "apple": 
		fmt.Println("This is an apple") 
	case "banana":
		fmt.Println("This is a banana")
	default: 
		fmt.Println("Unknown fruit")
	}

	// check multiple conditions in a switch statement
	day := "Monday"

	switch day {
		case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
			fmt.Println("It's a weekday")
		case "Saturday", "Sunday":
			fmt.Println("It's the weekend")
		default: 
			fmt.Println("Unknown day")
	}

	// we can use expressions in switch cases
	number := 15

	switch {
	case number < 10:
		fmt.Println("Number is less than 10");
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19");
	default: 
		fmt.Println("Number is either 20 or greater")
	}

	// using fallthrough to execute the next case
	num := 2

	switch {
	case num > 1: 
		fmt.Println("Greater than 1")
		fallthrough // this will execute the next case as well
	case num == 2:
		fmt.Println("Equal to 2")	
	default: 
		fmt.Println("Not 2")
	}

	checkType(10)
	checkType(3.14)
	checkType("Hello, Go!")
	checkType(true)
}

func checkType (x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("x is an integer")
	case float64: 
		fmt.Println("x is a float")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Println("x is of an unknown type")
	}
}