package main

import "fmt"

func main() {
	// basic if-else conditional statements
	// evaluate conditions
	if num := 10; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// basic switch-case statements
	// comparing the input with exact conditions
	var num int = 10
	switch num {
	case 1:
		fmt.Println("first")
	case 2:
		fmt.Println("second")
	default:
		fmt.Println("doesn't matter")
	}
}
