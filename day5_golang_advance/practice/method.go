package main

import (
	"fmt"
	"strings"
)

// define new struct
type student struct {
	name  string
	grade int
}

// define sayHello method using student struct as an argument
// same with function, but add arguments parameter
func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

// define getName method with integer input
func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

// define method with pointer object
func (s *student) changeName(name string) {
	fmt.Println("---> on ChangeName2, name changed to", name)
	s.name = name
}

// call method
// method will latch into the object, we can access method
func main() {
	var s1 = student{"john wick", 21}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan :", name)

	s1.changeName("ethan hunt")
	fmt.Println("s1 after changeName", s1.name)

	// still can access the object property direcly
	fmt.Println("=========================")
	fmt.Println(s1.name)
	fmt.Println(s1.grade)

}
