package intermediate

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

// Method  with value reciever
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer reciever
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func main() {

	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("Area of rectangle:", area)

	// pointer receiver
	rect.Scale(2)
	area = rect.Area()

	fmt.Println("Area of rectangle after scale:", area)

	num := MyInt(-5)
	num1 := MyInt(9)
	fmt.Println("Is num positive?", num.IsPositive())
	fmt.Println("Is num1 positive?", num1.IsPositive())

	fmt.Println(num.welcomeMessage())

	s := Shape{Rectangle: Rectangle{length: 10, width: 9}}
	fmt.Println(s.Area())
}

type MyInt int

// Method on user defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

func (MyInt) welcomeMessage() string {
	return "Welcome to my int type"
}
