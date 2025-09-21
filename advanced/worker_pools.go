package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

// simulate processing of tickets
func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of personID %d with total cost of %d\n", req.numTickets, req.personID, req.cost)

		time.Sleep(time.Second)
		results <- req.personID
	}
}

func workerPools() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	// start ticket processor worker
	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	// send requests
	for i := range numRequests {

		ticketRequests <- ticketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: (i + 1) * price}
	}

	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for %d processed successfully!\n", <-ticketResults)
	}
}

// // first thing it to create a worker function
// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing task %d\n", id, task)
// 		// simulate some work
// 		time.Sleep(time.Second)
// 		results <- task * 2
// 	}
// }

// ============== BASIC WORKER POOL PATTERN =================

// func main() {

// 	numWorkers := 3
// 	numJobs := 10
// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	// create workers
// 	for i := range numWorkers {
// 		go worker(i, tasks, results)
// 	}

// 	// send values  to the task channel
// 	for i := range numJobs {
// 		tasks <- i
// 	}

// 	close(tasks)

// 	// collect the results
// 	for range numJobs {
// 		result := <-results
// 		fmt.Println("Result:", result)
// 	}

// }
