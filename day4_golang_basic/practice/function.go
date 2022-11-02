package main

import "fmt"

// function without return values
func start() {
	fmt.Println("Hello World")
}

// define a function
// with integer input and output
func addition(a int, b int) int {
	return a + b
}

// function with multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	start()

	result := addition(1, 2)
	fmt.Println("1 + 2 =", result)

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
