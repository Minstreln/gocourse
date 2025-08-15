package basics

import "fmt"


func main() {

	// age := 25
	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// } 

	temperature := 25
	if temperature >= 30 {
		fmt.Println("It's a hot day.")
	} else {
		fmt.Println("It's a cool day.")
	}

	score := 85
	if score >= 90 {
		fmt.Println("You got an A.")
	} else if score >= 80 {
		fmt.Println("You got a B.")
	} else if score >= 70 {
		fmt.Println("You got a C.")
	} else {
		fmt.Println("You need to improve")
	}

	num := 15
	if num % 2 == 0 {
		if num % 3 == 0 {
			fmt.Println("Number is divisible by both 2 and 3")
		} else {
			fmt.Println("Number is divisible by 2 but not by 3")
		}
	} else {
		fmt.Println("Number is not divisible by 2")
	}

	// check for multiple conditions
	if 10 % 2 == 0 || 5 % 2 == 0 {
		fmt.Println("Either 10 or 4 are even")
	}

	if 10 % 2 == 0 && 5 % 2 == 0 {
		fmt.Println("both 10 or 4 are even")
	}
}