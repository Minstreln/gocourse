package basics

import (
	"fmt"
	"slices"
)


func main() {

	// var sliceName[]elementType

	// var numbers []int
	// var numbers1 = []int{1,2,3}

	// number2 := []int{9,8,7}

	// slice := make([]int, 5)

	a := [5]int{1,2,3,4,5}
	slice1 := a[1:4]

	fmt.Println(slice1)

	// append more elements to a slice
	slice1 = append(slice1, 6, 7)
	fmt.Println("Slice 1: ", slice1)

	// copy a slices
	sliceCopy :=  make([]int, len(slice1))
	copy(sliceCopy, slice1)

	fmt.Println("Slice copy: ", sliceCopy)

	// nil slice
	// var nilSlice []int

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	fmt.Println("Element at index 3: ", slice1[3])

	// slice1[3] = 50
	// fmt.Println("Element at index 3: ", slice1[3])

	// comparing two slices
	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("Slice1 is equal to sliceCopy")
	}

	// multi dimensional , two dimensional slice
	twoD := make([][]int, 3)
	for i := 0; i<3; i++ {
		innerLen := i + 1;
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i +j
			fmt.Printf("Adding %d in outer slice at index %d and inner slice index %d\n", i+j, i, j)
		}
	}

	fmt.Println(twoD)

	// slice operator - slice(low:high)
	slice2 := slice1[2:4]
	fmt.Println(slice2)

	fmt.Println("capacity of slice2 is ", cap(slice2))
	fmt.Println("Length of slice2 is ", len(slice2))

}

