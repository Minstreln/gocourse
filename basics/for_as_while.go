package basics

import "fmt"


func main() {

	i := 1
	for i <= 5 {
		fmt.Println("Iteration:", i)
		i++
	}

	// using for loop as while loop with break
	sum := 0
	for  {
		sum += 10
		fmt.Println("Current sum:", sum)
		if sum >= 50 {
			break
		}
	}

	// using for loop as while loop with continue
	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd number:", num)
		num++
	}

}