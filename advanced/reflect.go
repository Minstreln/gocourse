package advanced

import (
	"fmt"
	"reflect"
)

func main() {

	var x = 42
	v := reflect.ValueOf(x)
	t := v.Type()

	fmt.Println("Value:", v)
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
	fmt.Println("is int:", t.Kind() == reflect.Int)
	fmt.Println("is string:", t.Kind() == reflect.String)
	fmt.Println("Is zero:", v.IsZero())

}
