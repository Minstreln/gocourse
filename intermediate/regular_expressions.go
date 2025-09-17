package intermediate

import (
	"fmt"
	"regexp"
)

func main() {

	// fmt.Println("He said, \"I am great\"")
	// fmt.Println(`He said, "I am great"`)

	// compile a regex for an email address pattern
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// test strings
	email1 := "user@gmail.com"
	email2 := "invalid.email"

	// match
	fmt.Println("Email1: ", re.MatchString(email1))
	fmt.Println("Email2: ", re.MatchString(email2))

	// capturing groups
	// regex to capture date components
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	// test string
	date := "2025-12-31"

	// find submatches
	submatches := re.FindStringSubmatch(date)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	// taarget string

	str := "hello world, welcome to golang"

	re = regexp.MustCompile(`[aeiou]`)

	result := re.ReplaceAllString(str, "*")

	fmt.Println(result)

	// i - case insensitive match
	// g - global match
	// m - multi-line match
	// s - dotall mode (dot matches newline characters

	re = regexp.MustCompile(`(?i)go`)

	// test string

	text := "Golang is great. I love GO!"
	fmt.Println("Match: ", re.MatchString(text))

}
