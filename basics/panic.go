package basics

import "fmt"


func main() {

	// panic(interface{})

	// example of a valid input
	process(10)

	// example of an invalid input that will cause a panic
	process(-5)

}

func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("This is before panic")
		panic("Negative input not allowed")
	}

	fmt.Println("Processing input:", input)
}