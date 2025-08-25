package main

import "fmt"


func main() {

	// ...Elipses indicate a variadic function can accept zero or more arguments of a specified type.

	// fmt.Println("The sum of 1, 2, 3:", sum(1, 2, 3))
	// statement, total := sum("The sum of 1, 2, 3, 4, 5:", 1, 2, 3, 4, 5)
	statement, total := sum(1, 2, 3, 4, 5)
	fmt.Println(statement, total)

	// using slices with variadic functions
	numbers := []int{1, 2, 3, 4, 5, 9}

	statement2, total2 := sum(3, numbers...)
	fmt.Println("Statement:", statement2, "Total:", total2)
}


func sum(returnString int, nums ...int) (int, int) {
	total := 0 
	for _, v := range nums {
		total += v
	}

	return returnString, total
}