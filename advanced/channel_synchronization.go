package main

import (
	"fmt"
	"strconv"
	"time"
)

// func main() {

// 	done := make(chan struct{})

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		done <- struct{}{}
// 	}()
// 	<-done
// 	fmt.Println("Finished...")
// }

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		ch <- 9
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("Sent value to channel")
// 	}()

// 	value := <-ch
// 	fmt.Println(value)
// }

// =============== SYNCHRONIZING MULTIPLE GOROUTINES ===============
// func main() {
// 	numGoroutines := 3
// 	done := make(chan int, 3)

// 	for i := range numGoroutines {
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working...\n", id)
// 			time.Sleep(time.Second)
// 			done <- id
// 		}(i)
// 	}

// 	for range numGoroutines {
// 		<-done // wait for each goroutine to finish

// 	}

// 	fmt.Println("All goroutines finished.")

// }

// ============= SYNCHRONIZING DATA EXCHANGE =============
func synchronization() {

	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- "hello " + strconv.Itoa(i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()

	for value := range data {
		fmt.Println("Recieved Value:", value, ":", time.Now())
	}

}
