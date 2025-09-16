package intermediate

import (
	"fmt"
	"strconv"
)

func main() {
	str := "Hello, Go!"

	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World!"
	resutlt := str1 + " " + str2
	fmt.Println(resutlt)

	fmt.Println(str[0])

	// string conversion
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(str3)
}
