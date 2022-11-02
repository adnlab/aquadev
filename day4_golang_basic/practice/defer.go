package main

import "fmt"

func main() {
	// delay print to the end of execution
	defer fmt.Println("John")
	fmt.Println("Who are you?")
}
