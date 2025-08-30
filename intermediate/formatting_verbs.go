package intermediate

import "fmt"

func main() {

	// --- General formatting verbs in Go
	// %v  - prints the value in a default format
	// %#v - prints the value in Go syntax representation
	// %T  - prints the type of the value
	// %%  - prints a literal percent sign

	i := 15.5
	string := "Hello, Go!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	fmt.Printf("%v\n", string)
	fmt.Printf("%#v\n", string)
	fmt.Printf("%T\n", string)

	// --- Integer formatting verbs
	// %b  - base 2
	// %d - base 10
	// %+d - base 10 with sign
	// %o  - base 8
	// %#o - base 8 with leading zero
	// %x  - base 16, with a-f lowercase
	// %X  - base 16, with A-F uppercase
	// %#x - base 16, with leading 0x
	// %4d - pad with spaces (width 4, right justified)
	// %-4d - pad with spaces (width 4, left justified)
	// %04d - pad with zeros (width 4, right justified)

	var num int = 18

	fmt.Printf("Base 2: %b\n", num)
	fmt.Printf("base 10: %d\n", num)
	fmt.Printf("base 10 with sign: %+d\n", num)
	fmt.Printf("base 8: %o\n", num)
	fmt.Printf("base 8 with leading zero: %O\n", num)
	fmt.Printf("base 16, with a-f lowercase: %x\n", num)
	fmt.Printf("base 16, with A-F uppercase: %X\n", num)
	fmt.Printf("base 16, with leading 0x: %#x\n", num)
	fmt.Printf("pad with spaces (width 4, right justified): %4d\n", num)
	fmt.Printf("pad with spaces (width 4, left justified): %-4d\n", num)
	fmt.Printf("pad with zeros (width 4, right justified): %04d\n", num)

	// --- string formatting verbs
	// %s  - prints the value as a plain string
	// %q  - prints the value as a double-quoted string
	// %8s  - pring the value as a plain string, pad with spaces (width 8, right justified)
	// %-8s - prints the value as a plain string, pad with spaces (width 8, left justified)
	// %x - prints the value as a hex dump of byte values
	// % x - prints the value as a hex dump of byte values (uppercase letters)

	txt := "World"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	// --- Boolean formatting verbs
	// %t - prints the value as true or false (same as using %v)

	t := true
	f := false

	fmt.Printf("%t\n", t)
	fmt.Printf("%v\n", t)
	fmt.Printf("%t\n", f)

	//  ---  float formatting verbs
	// verb description
	// %e - scientific notation (e.g., 1.234456e+06)
	// %f - decimal point but no exponent (e.g., 1234456.234456)
	// %.2f - decimal point with 2 digits after the point (e.g., 1234456.23)
	// %6.2f - width 6, 2 digits after the point, pad with spaces (e.g., "1234.56")
	// %g - exponent as needed only necessary digits (e.g., 1234456 or 1.234456e+06)

	flt := 9.18

	fmt.Printf("%e\n", flt)
	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f\n", flt)
	fmt.Printf("%6.2f\n", flt)
	fmt.Printf("%g\n", flt)

}
