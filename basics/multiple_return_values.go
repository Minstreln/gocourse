package basics

import (
	"errors"
	"fmt"
)


func main() {

	q, r := divide(10, 3)
	fmt.Printf("Qoutient: %d. Remainder: %d\n", q, r)

	// send errors
	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

}

func divide(a, b int) (int, int) {
	qoutient := a / b
	remainder := a % b
	return qoutient, remainder
}

// returning error
func compare(a, b int) (string, error) {
	if (a > b ) {
		return "a is greater than b", nil
	} else if (b > a) {
		return "b is greater than a", nil
	} else {
		return "", errors.New("unable to compare which is greater")
	}
}