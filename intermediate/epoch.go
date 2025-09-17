package intermediate

import (
	"fmt"
	"time"
)

func main() {

	// 00:00:00 UTC jan 1, 1970

	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current unix time:", unixTime)

	t := time.Unix(unixTime, 0)
	fmt.Println("Converted time:", t)
}
