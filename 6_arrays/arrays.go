package main

import (
	"fmt"
	"slices"
)

func main() {

	// Default Values
	// Int -> 0, Bool -> false, String -> ""

	// Integer array
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	// String array
	var names [4]string = [4]string{"Devi", "Sri", "Sai"}

	for i, name := range names {
		println(i, name)
	}

	nums := [3]int{1, 2, 3}
	for i, num := range nums {
		println(i, num)
	}

	// 2D arrays
	two_d_arrays := [][]int{{1, 2}, {3, 4}, {5, 6}}
	for i, row := range two_d_arrays {
		for j, val := range row {
			println(i, j, val)
		}
	}

	// - Fixed size, cannot be resized
	// - Constant time access, O(1)
	// - Memory efficient, as elements are stored contiguously in memory

	// Slices
	// - Dynamic size, can be resized
	// - Built on top of arrays, provides more functionality
	// - Can be created using make() function or by slicing an existing array

	slice := []int{1, 2, 3, 4, 5} // No need to add the length
	slice = append(slice, 6)      // Adding an element to the slice

	print(len(slice)) // Length of the slice
	print(cap(slice)) // Capacity of the slice
	// capacity -> Maximum number of elements the slice can hold before it needs to be resized

	// Slices are more flexible than arrays and are commonly used in Go for working with collections of data.

	// Uninitialized slice is nil and has a length and capacity of 0
	var uninitializedSlice []int
	println(uninitializedSlice == nil) // true
	println(len(uninitializedSlice))   // 0
	println(cap(uninitializedSlice))   // 0

	// You can also create a slice using the make() function
	makeSlice := make([]int, 3, 5) // Creates a slice of length 3 and capacity 5
	println(len(makeSlice))        // 3
	println(cap(makeSlice))        // 5

	println(makeSlice)               // [0 0 0]
	makeSlice = append(makeSlice, 6) // Appending an element to the slice
	for i, v := range makeSlice {
		println(i, v)
	}
	println(len(makeSlice)) // 4 (length is increased by 1 after appending)

	println(cap(makeSlice)) // 5 (capacity remains the same until it needs to be resized)

	// Appending more elements to exceed the capacity
	makeSlice = append(makeSlice, 7, 8)
	println(len(makeSlice)) // 8 (length is increased by 4 after appending)

	println(cap(makeSlice)) // 10 (capacity is doubled when the slice needs to be resized)

	// Slice operator
	arr2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr2[:3]) // [1 2 3]

	// Slices package

	fmt.Println("Using Slices Package")

	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}

	fmt.Println(slices.Equal(slice1, slice2)) // true
}
