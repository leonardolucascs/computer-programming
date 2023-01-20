// array_slices.go

package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not inlcuded!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	//slice1[] = {2,3,4}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Print("\n\n")
	// more info
	fmt.Printf("The lenght of arr1 is %d\n", len(arr1))
	fmt.Printf("The lenght of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", len(slice1))

	fmt.Print("\n\n")
	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The lenght of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}
