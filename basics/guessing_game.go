package basics

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// welcome message to the user
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100.")
	fmt.Println("Try to guess it!")

	var guess int 
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		// check if the guess is correct
		if guess == target {
			fmt.Println("Congratulations! You guessed the number!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number.")
		} else {
			fmt.Println("Too high! Try guessing a lower number.")
		}
	}

}