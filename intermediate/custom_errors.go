package intermediate

import (
	"errors"
	"fmt"
)

func main() {

	if err := doSomething(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Operation successful")

}

type customError struct {
	code    int
	message string
	er      error
}

// error returns the error message Implementing Error() method of error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v", e.code, e.message, e.er)
}

// function that return a custom errror
// func doSomething() error {
// 	return &customError{
// 		code:    500,
// 		message: "Internal Server Error",
// 	}
// }

func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			er:      err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("An unexpected error occurred")
}
