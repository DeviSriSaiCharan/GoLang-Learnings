package main

import "fmt"

func main() {
	const pi = 3.14

	// We can declare multiple constants in a block
	const (
		port = 8080
		host = "localhost"
	)

	fmt.Println("Value of pi: ", pi)
	fmt.Println("Port: ", port)
	fmt.Println("Host: ", host)
}
