package main

import (
	"fmt"
	"math/rand"
)

// Since x and y are both of type int, we can omit the type for x and just write it once after y.
func add(x, y int) int {
	return x + y
}

func generateNumber() int {
	return rand.Intn(100)
}

// This function takes two function as arguments, that function type is func() int, which means it takes no arguments and returns an int.
func subtract(f1, f2 func() int) int {
	return f1() - f2()
}

// This function returns another function that prints the result. The inner function is a closure that captures the result variable from the outer function.
func printResult(result int) func() {
	return func() {
		fmt.Printf("The result is: %d\n", result)
	}
}

func multiply(x, y int) {
	printer := printResult(x * y)
	printer()
}

func main() {
	fmt.Println(add(5, 3))

	subtractResult := subtract(generateNumber, generateNumber)
	fmt.Printf("The result of subtraction is: %d\n", subtractResult)

	multiply(4, 7)
}
