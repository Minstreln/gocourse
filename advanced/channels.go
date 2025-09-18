package advanced

import "fmt"

func main() {

	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
		greeting <- "World"
		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}

	}()

	reciever := <-greeting
	fmt.Println(reciever)
	reciever = <-greeting
	fmt.Println(reciever)

	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}
}
