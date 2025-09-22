package advanced

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	// cmd := exec.Command("echo", "Hello world!")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Output:", string(output))

	cmd := exec.Command("grep", "foo")

	// set input for out command
	cmd.Stdin = strings.NewReader("foo\nbar\nbaz\n")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("output:", string(output))
}
