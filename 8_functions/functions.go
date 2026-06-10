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

func returnsMultipleValues() (int, string) {
	// Write something funny here
	return rand.Intn(100), "Here's a random number for you!"
}

// Can take a variable number of arguments of any type, and it will print each argument with its index.
func variadicFunction(random ...any) {
	for i, v := range random {
		fmt.Printf("Argument %d: %v\n", i, v)
	}
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(add(5, 3))

	subtractResult := subtract(generateNumber, generateNumber)
	fmt.Printf("The result of subtraction is: %d\n", subtractResult)

	multiply(4, 7)

	randomNumber, message := returnsMultipleValues()
	fmt.Printf("%s The random number is: %d\n", message, randomNumber)

	variadicFunction("Hello, world!", 42, true)

	nums := []int{1, 2, 3, 4, 5}

	// The ... operator is used to unpack the slice nums into individual arguments for the sum function.
	totalSum := sum(nums...)
	fmt.Printf("The total sum is: %d\n", totalSum)
}
