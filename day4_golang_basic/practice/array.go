package main

import "fmt"

func main() {
	// declare an empty array with defined range
	var numbers [4]int
	fmt.Println("emp:", numbers)

	// assigning value to elements of array
	numbers[0] = 3
	numbers[1] = 2
	numbers[2] = 1

	// accessing array values
	fmt.Println(numbers[0], numbers[1], numbers[2], numbers[3])
}
