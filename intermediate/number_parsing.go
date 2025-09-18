package intermediate

import (
	"fmt"
	"strconv"
)

func main() {

	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value", err)
	}

	fmt.Println(num)
	fmt.Println(num + 1)
}
