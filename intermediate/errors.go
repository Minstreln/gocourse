package intermediate

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math error: square root of negative number")
	}

	// compute square root
	return 1, nil

}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Data error: empty data")
	}
	// process data
	return nil
}

func main() {

	// result, err := sqrt(16)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)

	// result1, err1 := sqrt(-16)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// fmt.Println(result1)

	// data := []byte{}
	// err := process(data)
	// if err != nil {
	// if err := process(data); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Data processed successfully")

	// error interface  of custom error
	// if err1 := eprocess(); err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	if err := readData(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read successfully")
}

// custom errors
type myError struct {
	message string
}

func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eprocess() error {
	return &myError{"Custom error occurred"}
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData failed: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("Config Error")
}
