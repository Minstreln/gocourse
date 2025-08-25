package basics

import "fmt"


func main() {

	// func <name>(<parameters>) <return_type> { <body> }

	sum := add(1, 2)
	fmt.Println(sum)
	fmt.Println(add(2, 3))

	// making functions without eplicitly giving it a name(anonymous function)
	greet := func() {
		fmt.Println("Hello, anonymous function!")	
	}

	greet()

	// functions as types
	operation := add

	result  := operation(3, 4)
	fmt.Println("Result of operation:", result)

	// passing functions as arguments
	result2 := applyOperation(5, 3, add)
	fmt.Println("Result of applyOperation:", result2)

	// returning and using functions
	multiplyBy2 := createMultiplier(2)
	fmt.Println("5 multiplied by 2 is:", multiplyBy2(5)) 

}

func add(a, b int) int {
	return a + b
}

// function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}