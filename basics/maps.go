package basics

import (
	"fmt"
	"maps"
)


func main() {

	// var mapVariable map[keyType]valueType
	// using make - mapVariable = make(map[keyType]valueType)
	
	// using map literal
	// mapVariable = map[keyType]valueType {
	// 	key1: value1,
	// 	key2: value2,
	// }

	myMap := make(map[string]int)

	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	myMap["code"] = 18

	fmt.Println(myMap)
	// fmt.Println(myMap["key1"])
	// fmt.Println(myMap["key"])

	// delete a key value pair
	// delete(myMap, "key1")
	// fmt.Println(myMap)

	// completely remove all values
	// clear(myMap)
	// fmt.Println(myMap)

	_, ok := myMap["key1"]
	if ok {
		fmt.Println("A value exists with key1")
	} else {
		fmt.Println("A value does not exists with key1")
	}

	// fmt.Println(value)
	// fmt.Println("Is a value associated with key1:", ok)

	// initialize map in a different way
	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	myMap3 := map[string]int{"a": 1, "b": 2}

	// compare maps
	if maps.Equal(myMap2, myMap3) {
		fmt.Println("myMap2 and myMap3 are equal")
	}

	for k, v := range myMap3 {
		fmt.Println(k, v)
	}

	// the 0 values of a map is nil
	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value")
	} else {
		fmt.Println("The map is not initialized to nil value")
	}

	val := myMap4["key"]
	fmt.Println(val) // This will print an empty string since myMap4 is nil

	// myMap4["key"] = "value" // This will cause a panic since myMap4 is nil
	// fmt.Println(myMap4) // This line will not be reached due to the panic

	// if we want to enter a key into the map initialised to nil, we use make
	myMap4 = make(map[string]string)
	myMap4["key"] = "value"
	fmt.Println(myMap4)

	// getting the length of map
	fmt.Println("myMap4 length is:", len(myMap4))

	// multi dimensional map (nested maps)
	myMap5 := make(map[string]map[string]string)
	
	myMap5["map1"] = myMap4
	fmt.Println(myMap5)

}