package basics

import "fmt"

func main() {
	// simple iteration over a range
	for i := 1;  i <= 5; i++ {
		 fmt.Println(i)
	} 

	// iterate over a slice
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value:%d\n", index, value)
	}

	// using breaak and continue statements
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Println("Odd number:", i)
		if i == 5 {
			break // exit the loop when i is 5
		}
	}

	// asterik pyramid pattern
	rows := 5

	// outer loop
	for i := 1; i <= rows; i++ {
		// inner loop for printing spaces
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		// inner loop for printing stars
		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println() // move to the next line after each row
	}

	for i := range 10 {
		fmt.Println(10 - i)
	}
	fmt.Println("We have a lift off!")
}