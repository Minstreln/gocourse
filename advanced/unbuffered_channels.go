package advanced

import "fmt"

func advanced() {

	// buffered means a channel with storage - buffered channels allows asynchronous communication
	ch := make(chan int)
	go func() {
		ch <- 1
	}()

	receiver := <-ch
	fmt.Println(receiver)

}
