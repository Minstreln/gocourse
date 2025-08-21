package basics

import (
	"fmt"
	"math"
)

func main() {
	// variables declaration
	var a, b = 10, 3
	var result int

	result = a + b // Addition
	fmt.Println("Addition:", result)

	result = a - b // subtraction
	fmt.Println("Subtraction:", result)

	result = a * b // Multiplication
	fmt.Println("Multiplication:", result)

	result = a / b // Division - in Go, division of two integers results in an integer
	fmt.Println("Division:", result)

	result = a % b // Modulus Remainder
	fmt.Println("Modulus:", result)

	const p float64 = 22 / 7.0
	fmt.Println("Constant p:", p)

	// overflow example with signed integer
	var maxInt int64 = 9223372036854775807 // Max value for int64
	fmt.Println("Max Int64:", maxInt)  

	maxInt = maxInt + 1 // this will overflow
	fmt.Println(maxInt)

	// overflow example with unsigned integer
	var uMaxInt uint64 = 18446744073709551615 // max value for uint64
	fmt.Println("Max Uint64:", uMaxInt)

	uMaxInt = uMaxInt + 1 // this will overflow
	fmt.Println(uMaxInt) // will print 0 due to overflow

	// example of underflow
	var smallFloat float64 = 1.0e-323 // Min value for float64
	fmt.Println(smallFloat)
	
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}



