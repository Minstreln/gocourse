package intermediate

import (
	"fmt"
	"math/rand"
)

func main() {

	// fmt.Println(rand.Intn(6) + 5)

	val := rand.New(rand.NewSource(42))

	fmt.Println(val.Intn(6) + 5)
}
