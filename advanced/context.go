package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled", ctx.Err())
			return
		default:
			fmt.Println("Doing work...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func cont() {
	rootCtx := context.Background()
	ctx, cancel := context.WithTimeout(rootCtx, 2*time.Second) // timer of the context starts here
	defer cancel()

	ctx = context.WithValue(ctx, "requestID", "skjdnksd123")

	go doWork(ctx)

	time.Sleep(3 * time.Second)

	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("Request ID:", requestID)

	} else {
		fmt.Println("No Request ID found")
	}

	logWithContext(ctx, "This is a test log messaage")

}

func logWithContext(ctx context.Context, message string) {
	requestIdValue := ctx.Value("requestID")
	log.Printf("Request ID: %v - %v", requestIdValue, message)
}

// func checkEvenOdd(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done():
// 		return "Operation cancelled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is Even", num)
// 		} else {
// 			return fmt.Sprintf("%d is Odd", num)
// 		}

// 	}
// }

// func main() {
// 	ctx := context.TODO()

// 	result := checkEvenOdd(ctx, 5)
// 	fmt.Println("Result with context.TODO():", result)

// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	result = checkEvenOdd(ctx, 10)
// 	fmt.Println("Result with context.WithTimeout():", result)

// 	time.Sleep(2 * time.Second)
// 	result = checkEvenOdd(ctx, 15)
// 	fmt.Println("Result after timeout:", result)

// }

// =======  DIFFERENCE BETWEEN CONTEXT.TODO() AND CONTEXT.BACKGROUND()  =======
// func main() {

// 	todoContext := context.TODO()
// 	contextBkg := context.Background()

// 	ctx := context.WithValue(todoContext, "name", "john")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name"))

// 	ctx1 := context.WithValue(contextBkg, "city", "New York")
// 	fmt.Println(ctx1)
// 	fmt.Println(ctx1.Value("city"))

// }
