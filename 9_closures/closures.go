package main

import "fmt"

// If a function that references the variables that is defined outside of the function, it is called a closure.
// A closure allows the function to access and manipulate the variables that are defined outside of its scope, even after the outer function has finished executing.
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := makeCounter()
	fmt.Println(counter()) // Output: 1
	fmt.Println(counter()) // Output: 2
	fmt.Println(counter()) // Output: 3

	counter2 := makeCounter()
	fmt.Println(counter2()) // Output: 1
	fmt.Println(counter2()) // Output: 2
}
