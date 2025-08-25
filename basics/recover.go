package basics

import "fmt"

func main() {

	process()
	fmt.Println("Returned from process...")

}

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("processing...")
	panic("something went wrong")

}
