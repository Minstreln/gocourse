package intermediate

import (
	"fmt"
	"time"
)

func main() {

	// Mon Jan 2 15:04:05 MST 2006
	layout := "2006-01-02T15:04:05Z07:00"
	str := "2023-10-05T14:30:00+00:00"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	fmt.Println(t)

	str1 := "Jul 03, 2024 03:18 PM"
	layout1 := "jan 02, 2006 03:04 PM"

	t1, err := time.Parse(layout1, str1)
	fmt.Println(t1)

}
