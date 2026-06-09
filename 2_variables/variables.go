package main

import "fmt"

func main() {

	// Declaring a variable with a specific type and value
	var name string = "Charan"

	// Declaring a variable with type inference (the type is inferred from the assigned value)
	var age = 22

	// Shorthand variable declaration (only works inside functions)
	// We cantt use this outside of a function
	// We can only use this when we are declaring and initializing a variable at the same time
	isStudnet := false

	// In this case we can use shorthand declaration
	var city string

	// ... after some time

	city = "Vijayawada"

	fmt.Println("Name: ", name)

	fmt.Println("Age: ", age)

	fmt.Println("Is Student: ", isStudnet)

	fmt.Println("City: ", city)
}
