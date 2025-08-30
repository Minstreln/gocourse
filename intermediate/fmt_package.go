package intermediate

import "fmt"

func main() {

	// Printing functions in Go's fmt package
	fmt.Print("Hello ")
	fmt.Print("World!")
	fmt.Print(12, 456)

	fmt.Println("Hello ")
	fmt.Println("World!")
	fmt.Println(12, 456)

	name := "John"
	age := 25

	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	// formatting functions in Go's fmt package
	s := fmt.Sprint("Hello ", "World!", 12, 456)
	fmt.Print(s)

	s = fmt.Sprintln("Hello", "World!", 12, 456)
	fmt.Print(s)
	fmt.Print(s)

	sf := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(sf)
	fmt.Println(sf)

	// scanning functions
	// var nameInput string
	// var ageNum int

	// fmt.Print("Enter your name and age: ")
	// fmt.Scan(&nameInput, &ageNum)
	// fmt.Printf("Name: %s, Age: %d\n", nameInput, ageNum)

	// scanln function requires pressing Enter after input and it has to be one input ant it stops scanning at a new line

	// var nameInput string
	// var ageNum int

	// fmt.Print("Enter your name and age: ")
	// fmt.Scanln(&nameInput, &ageNum)
	// fmt.Printf("Name: %s, Age: %d\n", nameInput, ageNum)

	// scanf

	// var nameInput string
	// var ageNum int

	// fmt.Print("Enter your name and age: ")
	// fmt.Scanf("%s %d", &nameInput, &ageNum)
	// fmt.Printf("Name: %s, Age: %d\n", nameInput, ageNum)

	// Error formatting functions

	err := checkAge(15)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is to young to drive", age)
	}

	return nil
}
