package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1

	// Switch dont need break statement, it automatically breaks after executing a case
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Default case")
	}

	// switch with multiple cases
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// Here interface{} is a type that can hold any value, and we use type assertion to determine the actual type of the value at runtime.
	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("I am an integer: %d\n", v)
		case string:
			fmt.Printf("I am a string: %s\n", v)
		default:
			fmt.Printf("I am of a different type: %T\n", v)
		}
	}

	whoAmI(42)
	whoAmI("Hello, Go!")
	whoAmI(3.14)
}
