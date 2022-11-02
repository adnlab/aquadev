package main

import (
	"fmt"
)

func main() {
	// declare and initiate a new variable with integer data type
	var a int = 1
	fmt.Println(a)

	// separate variable declaration and initiation
	var b int
	b = 2
	fmt.Println(b)

	// define a new variable without data types
	var c = 3
	fmt.Println(c)

	// another way to declare and initiate a new variable
	d := 4
	fmt.Println(d)

	// multi variables declaration
	var first, second, third string
	first, second, third = "one", "two", "three"

	fourth, sixth := "four", "six"

	fmt.Println(fourth, sixth)

	// using trash variable for unused variable
	_, _, _ = first, second, third

	// define a constant variable
	const pi = 3.14159265359
	fmt.Println(pi)
}
