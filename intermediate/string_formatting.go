package intermediate

import "fmt"

func main() {

	num := 23432343
	fmt.Printf("%05d\n", num)

	// text alignment
	message := "Hello"
	fmt.Printf("|%10s|\n", message)  // right align
	fmt.Printf("|%-10s|\n", message) // left align

	// string interpolation
	message1 := "World \nWorld!"
	message2 := `World \nWorld!`

	fmt.Println(message1)
	fmt.Println(message2)

	sqlQuery := `SELECT * FROM  users WHERE age > 30`
	fmt.Println(sqlQuery)
}
