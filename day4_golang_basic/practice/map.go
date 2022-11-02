package main

import "fmt"

func main() {
	// declare and initiate a map separately
	var m map[string]int
	m = map[string]int{}
	fmt.Println("map:", m)

	// assign a value to an empty map
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// accessing a map value
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// return a map length
	fmt.Println("len:", len(m))

	// delete a map value
	delete(m, "k2")
	fmt.Println("map:", m)

	fmt.Println("++++++++++++++++++++")

	// another way to declare an empty map, using make
	n := make(map[string]int)
	fmt.Println("map:", n)

	// declare and initiate map
	p := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", p)

}
