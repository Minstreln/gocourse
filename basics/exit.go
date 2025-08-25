package basics

import (
	"fmt"
	"os"
)

func main() {

	// avoid deferred actions when exiting
	defer fmt.Println("deferred function called")

	fmt.Println("starting main function")

	// exit with a status code of 1
	os.Exit(1)

	// this will never be printed
	fmt.Println("End of main function")

}
