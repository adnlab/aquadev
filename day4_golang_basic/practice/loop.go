package main

import "fmt"

func main() {
	// will loop until the condition are met
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// will loop for a specific defined range
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// iterate over a data types

	// slice
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// array

	// map

	// struct

}
