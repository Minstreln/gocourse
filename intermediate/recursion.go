package basics

import "fmt"

func main() {

	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

	fmt.Println(sumOfDigits(9))
	fmt.Println(sumOfDigits(12))
	fmt.Println(sumOfDigits(12345))

}

// function to calculate factorial of a number
func factorial(n int) int {
	// base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}

	// recursive case: factorial of n is n * factorial of (n-1)
	return n * factorial(n-1)
	// n * (n-1) * (n-2) * factorial(n-3) ... 1
}

// sum of digits
func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}

	return n%10 + sumOfDigits(n/10)
}
