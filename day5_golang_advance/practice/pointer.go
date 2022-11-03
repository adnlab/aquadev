package main

import "fmt"

func main() {
	// taking a normal variable
	var x int = 100

	// declaration of pointer variable
	var p *int

	// initialization of pointer
	p = &x

	// displaying the result
	fmt.Println("Value stored in x =", x)
	fmt.Println("Address of x =", &x)
	fmt.Println("Value stored in p =", p)
	fmt.Println("=======================")

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println(numberA)
	fmt.Println(&numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)

	fmt.Println("change numberA value")

	// can use either of these
	numberA = 5
	//*numberB = 5

	fmt.Println(numberA)
	fmt.Println(&numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)
}
