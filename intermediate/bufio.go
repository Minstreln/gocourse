package intermediate

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	reader := bufio.NewReader(strings.NewReader("Hello, bufio package!\n"))

	// reading byte slice
	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string:", err)
		return
	}

	fmt.Println("Read line:", line)
}
