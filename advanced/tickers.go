package main

import (
	"fmt"
	"time"
)

func ticker() {

	ticker := time.NewTicker(time.Second)

	for tick := range ticker.C {
		fmt.Println("Tick at: ", tick)
	}

}
