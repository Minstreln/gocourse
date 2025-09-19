package main

import (
	"fmt"
	"time"
)

// func main() {

// 	// =========== BLOCKING ON RECEIVE ONLY CHANNEL WHEN THE BUFFER IS EMPTY ===========
// 	// make(chan type, capacity)
// 	ch := make(chan int, 2)
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 	}()

// 	fmt.Println("Value: ", <-ch)
// 	fmt.Println("End of Program")

// }

func buffered() {

	// ========= BLOCKING ON SEND ONLY CHANNEL WHEN THE BUFFER IS FULL ========

	// make(chan type, capacity)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 // This will cause a deadlock since the buffer is full
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Recieved: ", <-ch)
	}()

	ch <- 3 // This will work since there is space in the buffer after the goroutine reads from it
	fmt.Println("Recieved: ", <-ch)
	fmt.Println("Recieved: ", <-ch)

	fmt.Println("Buffered channels")

}
