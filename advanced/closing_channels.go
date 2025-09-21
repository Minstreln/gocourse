package main

import "fmt"

// ========== RANGE OVER CLOSED CHANNEL ============
func closingchannels() {
	ch := make(chan int)

	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println("Recieved:", val)
	}
}

// =========== RECIEVING FROM A CLOSED CHANNEL =======
// func main() {
// 	ch := make(chan int)
// 	close(ch)

// 	val, ok := <-ch
// 	if !ok {
// 		fmt.Println("Channel is closed")
// 		return
// 	}
// 	fmt.Println(val)

// }

// ======= SIMPLE CLOSING CHANNEL =====
// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for value := range ch {
// 		fmt.Println(value)
// 	}

// }
