package main

import (
	"fmt"
)

func multiplexing() {
	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch)
	}()

	for range 1 {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Received from channel:", msg)

		default:
			fmt.Println("No communication")
		}
	}
}

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		close(ch)
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println("Received from channel:", msg)
// 	case <-time.After(3 * time.Second):
// 		fmt.Println("Timeout: No communication within 1 second")
// 	}
// }

// func main() {

// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {
// 		// time.Sleep(time.Second)
// 		ch1 <- 1
// 	}()

// 	go func() {
// 		// time.Sleep(time.Second)
// 		ch2 <- 1
// 	}()

// 	// wait for goroutines to send data
// 	time.Sleep(1 * time.Second)

// 	for range 2 {
// 		select {

// 		case msg := <-ch1:
// 			fmt.Println("Recieved from channel 1:", msg)

// 		case msg := <-ch2:
// 			fmt.Println("Recieved from channel 2:", msg)
// 		default:
// 			fmt.Println("No communication")
// 		}
// 	}

// }
