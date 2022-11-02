package main

import "fmt"

func main() {
	// define an empty slice with specific range
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// assign new value to defined slice
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("get:", s[2])
	fmt.Println("length:", len(s))

	// can't add value past the defined slice range
	// s[3] = "d"

	// define a flexible ranged slice
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits)
	fmt.Println("length:", len(fruits))
}
