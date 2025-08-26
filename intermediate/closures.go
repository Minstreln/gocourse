package intermediate

import "fmt"

func main() {

	sequence := adder()
	fmt.Println(sequence())

	subtractor := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	// using the closure subtractor
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(2))
	fmt.Println(subtractor(3))
	fmt.Println(subtractor(4))
	fmt.Println(subtractor(5))
}

func adder() func() int {
	i := 0
	fmt.Println("Previous value of i:", i)
	return func() int {
		i++
		fmt.Println("adding 1 to i:")
		return i
	}
}
