package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!" // \n is for a tab space
	message2 := "Hello, \rGo!" // \r is a carriage return
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of variable message is:", len(message))

	fmt.Println(message[0]) // prints ASCII value of first character

	// string concatenation
	greeting := "Hello "
	name := "Alice"
	msg := greeting + name
	fmt.Println(msg)

	// compare two strings
	string1 := "apple"
	string2 := "banana"
	fmt.Println(string1 < string2) // true because "apple" comes before "banana" lexicographically

	// string iteration
	for i, char := range message {
		fmt.Printf("Character at index %d is %c\n", i, char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	// string manipulation
	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	// Runes is n alias for int32 type, that represents an integer value of a Unicode code point

	// declaration of a rune, single quotes are used for runes
	var ch rune = 'a'
	fmt.Println(ch)
	fmt.Printf("Character: %c, Unicode code point: %U\n", ch, ch)

	// converting runes to strings
	cstr := string(ch)
	fmt.Println(cstr)

	// format verb to get the type of any variable
	fmt.Printf("type of cstr is: %T\n", cstr)

	const NIHONGO = "これは本です" // This is a book in Japanese
	fmt.Println(NIHONGO)

	// iterating over runes
	jhello := "こんにちは" // Hello in Japanese
	for _, runeValue := range jhello {

		fmt.Printf("%c\n", runeValue)
	}

	r := '🙂'
	fmt.Printf("Character: %c, Value: %v\n", r, r)

}
