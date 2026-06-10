package main

import (
	"fmt"
	"maps"
)

func main() {

	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println(m1)
	fmt.Println(m1["two"])
	fmt.Println(m1["zero"])

	val, isPresent := m1["zero"]
	fmt.Println(val, isPresent)

	// to delete a key from map
	delete(m1, "two")
	fmt.Println(m1)

	// to clear a map
	clear(m1)
	fmt.Println(m1)

	m2 := make(map[string]string)
	fmt.Println(m2)

	m2["name"] = "Devi Sri Sai Charan"
	m2["age"] = "22"
	fmt.Println(m2)

	m3 := map[string]string{
		"name": "Devi Sri Sai Charan",
		"age":  "22",
	}
	// There is a Maps package
	fmt.Println("Is Equal: ", maps.Equal(m2, m3))

}
