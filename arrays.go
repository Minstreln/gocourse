package main

import "fmt"

func main() {

	// var arrayName [size]elementType

	var numbers [5]int
	fmt.Println(numbers)
	
	numbers[4] = 20
	fmt.Println(numbers)

	// initializing an array with values

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("Fruits Array: ", fruits)

	fmt.Println("Third Fruit: ", fruits[2])

	// originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray
	// copiedArray[0] = 100

	// fmt.Println("Original array: ", originalArray)
	// fmt.Println("Copied array: ", copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index,", i, ":", numbers[i])
	}

	// use a range based iteration
	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// to discard a value we use the underscore(blank identifier)
	for _, v := range numbers {
		fmt.Printf("Value: %d\n", v)
	}

	a, _ := someFunction()
	fmt.Println(a)
	// fmt.Println(b)

	fmt.Println("The length of the numbers array is: ", len(numbers))

	// comparing  arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{10, 2, 3}

	fmt.Println(array1 == array2)

	// multi dimensional array
	var matrix [3][3]int = [3][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	fmt.Println("Matrix: ", matrix)

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println("Original array: ", originalArray)
	// fmt.Println("Copied array: ", *copiedArray)
}

func someFunction() (int, int) {
	return 1, 2
}