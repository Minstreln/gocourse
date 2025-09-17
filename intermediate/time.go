package intermediate

import (
	"fmt"
	"time"
)

func main() {

	// get current time
	fmt.Println(time.Now())

	// specific time
	specificTime := time.Date(2023, time.March, 14, 15, 9, 26, 0, time.UTC)
	fmt.Println(specificTime)

	// parse time
	parsedTime, _ := time.Parse("2006-01-02 15:04:05", "2023-03-14 15:09:26")
	parsedTime1, _ := time.Parse("06-01-02", "03-03-14")
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)

	// formatting time
	t := time.Now()
	fmt.Println("Formatted time:", t.Format("Monday 06-01-02 15:04:05"))

	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println("One day later:", oneDayLater)
	fmt.Println("One day later:", oneDayLater.Weekday())

	fmt.Println("Rounded time:", t.Round(time.Hour))

	// loc, _ := time.LoadLocation("Asia/Kolkata")
	// t = time.Date(2023, time.March, 14, 15, 9, 26, 0, time.UTC)

	// // convert tto this specific time zone
	// tLocal := t.In(loc)

	// // perform rounding
	// roundedTime := t.Round(time.Hour)
	// roundedTimeLocal := roundedTime.In(loc)

	// fmt.Println("Original time (UTC):", t)
	// fmt.Println("Original time (Local):", tLocal)
	// fmt.Println("Rounded time (UTC):", roundedTime)
	// fmt.Println("Rounded time (Local):", roundedTimeLocal)

	fmt.Println("Truncated Time:", t.Truncate(time.Hour))

	loc, _ := time.LoadLocation("America/New_York")

	// convert time to location
	tinNY := time.Now().In(loc)
	fmt.Println("Time in New York:", tinNY)

	t1 := time.Date(2023, time.March, 14, 15, 9, 26, 0, time.UTC)
	t2 := time.Date(2023, time.March, 15, 16, 10, 27, 0, time.UTC)

	duration := t2.Sub(t1)
	fmt.Println("Duration between t1 and t2:", duration)

}
