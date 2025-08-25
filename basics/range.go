package basics

import "fmt"


func main() {

	// use range over strings
	message := "Hello, World!";

	for i, v := range message {
		fmt.Println(i, v)
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}

	// for arrays, slices and strings, range iterates from the first element to the last, for maps it iterates over key value pairs but in an unspecified order and for channles it iterates over values received from the channel until it is closed.

	

}