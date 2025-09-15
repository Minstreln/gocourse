package intermediate

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person // Embedding Person struct - anonymous field
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

// you can override methods of embedded structs
func (e Employee) introduce() {
	fmt.Printf("Hi, I am %s, my employee ID is %s and I earn %.2f.\n", e.name, e.empId, e.salary)
}

func main() {

	emp := Employee{
		person: person{name: "Alice", age: 30},
		empId:  "E123",
		salary: 75000.00,
	}

	fmt.Println("name:", emp.name)     // Accessing embedded struct field directly without needing to do emp.person.name
	fmt.Println("age:", emp.age)       // Accessing embedded struct field directly
	fmt.Println("empId:", emp.empId)   // Accessing embedded struct field directly
	fmt.Println("salary:", emp.salary) // Accessing embedded struct field directly

	emp.introduce()
}
