package intermediate

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

// anonymous fields promote the fields of the embedded struct to the outer struct
type PhoneHomeCell struct {
	home string
	cell string
}

// embedding structs
type Address struct {
	city    string
	country string
}

// methods
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

// modifying structs using a pointer reciever
func (p *Person) incrementAgeBy1() {
	p.age++
}

func main() {

	// structs are composite data types that allow you to group together variables of different types under a single name similar to classes in other languages but they are more light weight and do not support inheritance

	p1 := Person{
		firstName: "John",
		lastName:  "Zhang",
		age:       30,
		address: Address{
			city:    "New York",
			country: "USA",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "123-456-7890",
			cell: "987-654-3210",
		},
	}

	p2 := Person{
		firstName: "Jane",
		age:       25,
	}

	p3 := Person{
		firstName: "Jane",
		age:       25,
	}

	// adding values to embedded struct
	// p2.address.city = "Los Angeles"
	// p2.address.country = "USA"

	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println(p1.fullName())
	fmt.Println(p1.address)
	fmt.Println(p2.address.city)
	fmt.Println(p1.cell)

	// comparing structs
	fmt.Println(p3 == p2)

	// Anonymous Structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "johndoe@gmail.com",
	}

	fmt.Println(user.username)
	fmt.Println("Before increment", p1.age)
	p1.incrementAgeBy1()
	fmt.Println("After increment", p1.age)
}
