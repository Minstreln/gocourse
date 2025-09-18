package intermediate

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("Hello, World!")

	// encode to base 64
	encode := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded:", encode)
}
