package main

import "fmt"

func rectEnTri(n int) {
	if n == 0 {
		fmt.Printf("Anda tidak bisa menggunakan n dengan nilai nol.")
	} else if n%2 == 1 {
		triShape(n)
	} else {
		rectShape(n)
	}
}

func triShape(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func rectShape(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func main() {
	fmt.Printf("Masukkan nilai n yang Anda inginkan: ")

	var number int
	fmt.Scanln(&number)

	rectEnTri(number)

}
