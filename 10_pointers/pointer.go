package main

import "fmt"

func increment(num int) {
	num++
}

func incrementPointer(num *int) {
	*num++
}

func main() {

	num := 5

	fmt.Println("Value of num: ", num)

	increment(num)

	fmt.Println("Value of num after increment: ", num)

	incrementPointer(&num)

	fmt.Println("Value of num after incrementPointer: ", num)

	fmt.Println("Address: ", &num)
}
