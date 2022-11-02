package main

import "fmt"

// define data "blueprint" with struct
type person struct {
	name string
	age  int
	job
}

// define a struct that will be embedded
type job struct {
	company string
	roles   string
}

func main() {
	// declare a variable with struct data types
	var p1 person

	// assign value to a struct variable
	p1.name = "John"
	p1.age = 42

	// accessing a struct variable
	fmt.Println("Name:", p1.name)
	fmt.Println("Age:", p1.age)

	fmt.Printf("\n")

	// define new struct object
	var p2 person
	p2.name = "Doe"
	p2.age = 30
	p2.company = "Madeup Company"
	p2.roles = "Software Engineer"

	fmt.Println("Name:", p2.name)
	fmt.Println("Age:", p2.age)
	fmt.Println("Company:", p2.company)
	fmt.Println("Job Roles:", p2.roles)

	fmt.Printf("\n")

	// define a slice using struct data types
	var allEmployee = []person{
		{name: "Ethan", age: 54},
		{name: "Anderson", age: 63},
	}

	// iterate over a slice, for objects value
	for _, employee := range allEmployee {
		fmt.Println(employee.name, "age is", employee.age)
	}

}
