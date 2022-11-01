package main

import "fmt"

type student struct {
	name string
}

func (s *student) sayHello(age int) string {
	return "hello, " + s.name + ", dengan umur: " + fmt.Sprint(age)
}

func main() {
	var x int = 5768

	var p *int

	p = &x

	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p)

	//

	var s1 = student{
		name: "Sir",
	}

	fmt.Println(s1.sayHello(15))

}
