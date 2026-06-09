package main

import "fmt"

// for -> only loop in go

func main() {
	// while loop
	i := 0

	for i < 3 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }

	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	// for range -> used to iterate over arrays, slices, maps, strings, etc.
	// -> Here 5 is exclusive, so it will print 0, 1, 2, 3, 4
	for i := range 5 {
		fmt.Println(i)
	}

	for i := range []string{"Go", "Python", "Java"} {
		fmt.Println(i)
	}

	for i, v := range []string{"Go", "Python", "Java"} {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	for _, v := range []string{"Go", "Python", "Java"} {
		fmt.Printf("Value: %s\n", v)
	}

	for k, v := range map[string]int{"Go": 1, "Python": 2, "Java": 3} {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

}
