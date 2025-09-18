package intermediate

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	// write data to file
	data := []byte("Hello, World!\nThis is a test file.\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data written to file successfully")

	file, err = os.Create("writeString.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	// write string to file
	_, err = file.WriteString("Hello, World!\nThis is a test file.\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("String written to file successfully")
}
