package main

import (
	"fmt"
	"time"
)

func timer() {
	timer1 := time.NewTimer(time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	select {
	case <-timer1.C:
		fmt.Println("Timer 1 expired")
	case <-timer2.C:
		fmt.Println("timer 2 expired")
	}
}

// func main() {
// 	timer := time.NewTimer(2 * time.Second)

// 	go func() {
// 		<-timer.C
// 		fmt.Println("Delayed operation executed")
// 	}()

// 	fmt.Println("Waiting...")
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("End of program")
// }

// func longRunningOperation() {
// 	for i := range 20 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// ========== TIMEOUT ========
// func main() {
// 	timeout := time.After(3 * time.Second)
// 	done := make(chan bool)

// 	go func() {
// 		longRunningOperation()
// 		done <- true
// 	}()

// 	select {
// 	case <-timeout:
// 		fmt.Println("Operation times out")
// 	case <-done:
// 		fmt.Println("Operation completed")
// 	}

// }

// func main() {

// 	fmt.Println("Starting app....")
// 	timer := time.NewTimer(2 * time.Second)
// 	fmt.Println("Waiting for timer.c")
// 	stopped := timer.Stop()
// 	if stopped {
// 		fmt.Println("Timer stopped")
// 	}

// 	timer.Reset(time.Second)
// 	fmt.Println("Waiting for timer reset")

// 	<-timer.C
// 	fmt.Println("Timer expired")

// }
