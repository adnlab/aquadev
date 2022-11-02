## What is Go Language

Go language is a new, general-purpose programming language that makes it easy to build simple, reliable, and efficient software.

[Golang](https://golang.org/) (atau biasa disebut dengan Go) adalah bahasa pemrograman yang dikembangkan di Google oleh [Robert](https://github.com/griesemer) [Griesemer](https://github.com/griesemer), [Rob Pike](https://en.wikipedia.org/wiki/Rob_Pike), dan [Ken Thompson](https://en.wikipedia.org/wiki/Ken_Thompson) pada tahun 2007 dan mulai diperkenalkan ke publik tahun 2009.

- Penciptaan bahasa Go didasari bahasa C dan C++, oleh karena itu gaya sintaksnya mirip

Company using Go:
- Gojek
- Tokopedia
- eFishery
- Google

## Initiate Go Project
- Installing Go language
- Initialize

```
mkdir <nama-project>
cd <nama-project>

go mod init <nama-project>
go run main.go
```

## Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

## Data Types

Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.
- Strings, which can be added together with `+`.
- Integers and floats.
- Booleans, with boolean operators as you’d expect.

## Variables

```go
var a int = 1

var b int
b = 2

//without data types
c := 3

//multi variables
var first, second, third string
first, second, third = "one", "two", "three"

fourth, sixth := "four", "six"

```

**Trash variable**
- to store unused variable
```go
a = 5
_ = a
```

**Constant**
- Constant value.
- Go supports _constants_ of character, string, boolean, and numeric values.
- Constant expressions perform arithmetic with arbitrary precision.

```go
const pi = 3.14159265359
```

## Condition

**If-Else**
- option based on condition
```go
if num := 9; num < 0 {
    fmt.Println(num, "is negative")
} else if num < 10 {
    fmt.Println(num, "has 1 digit")
} else {
    fmt.Println(num, "has multiple digits")
}
```

**Switch Case**
- specific value condition
```go
var num int = 2
switch num {
	case 1:
		fmt.Println("first")
	case 2:
		fmt.Println("second")
	default:
		fmt.Println("doesn't matter")
}
```

**Comparison operator**
several comparison operator:
- AND operator (&&), return true if both statements are true
- OR operator (||), return true if one of the statements are true
- NOT operator (!), reverse the result, return false if the result is true and vice-versa

## Loop
- `for` is Go’s only looping construct. Here are some basic types of `for` loops.
```go
i := 1
for i <= 3 {
    fmt.Println(i)
    i = i + 1
}

for j := 7; j <= 9; j++ {
    fmt.Println(j)
}
```

- `for` without a condition will loop repeatedly until you `break` out of the loop or `return` from the enclosing function.
- You can also `continue` to the next iteration of the loop.

```go
for i := 1; i <= 10; i++ {
	if i % 2 == 1 {
		continue
	}
	if i > 8 {
		break
	}
	fmt.Println("number", i)
}
```
- _range_ iterates over elements in a variety of data structures either its slice, array, map.
```go
nums := []int{2, 3, 4}
sum := 0
for _, num := range nums {
    sum += num
}
fmt.Println("sum:", sum)
```

## Function
- Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
```go
func addition(a int, b int) int {
	return a + b
}
```
- When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
- Called a function inside main() function
```go
func main() {
	result := addition(1, 2)
	fmt.Println("1+2=", result)
}
```
- Multiple return values
```go
func swap(x, y string) (string, string) {
	return y, x	
}
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

## Data Structure

**Array**
- In Go, an _array_ is a numbered sequence of elements of a specific length.
- The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for `int`s means `0`s.
```go
var numbers [4]int
fmt.Println("emp:", numbers)

numbers[0] = 3
numbers[1] = 2
numbers[2] = 1

fmt.Println(numbers[0], numbers[1], numbers[2], number[3])
```

**Slice**
- _Slices_ are an important data type in Go, giving a more powerful interface to sequences than arrays.
- Unlike arrays, slices are typed only by the elements they contain (not the number of elements). To create an empty slice with non-zero length, use the builtin `make`.
- We can set and get just like with arrays.
```go
s := make([]string, 3)
fmt.Println("emp:", s)

s[0] = "a"
s[1] = "b"
s[2] = "c"
fmt.Println("get:", s[2])
fmt.Println("length:", len(s))
```
- declare and initialize a variable for slice
```go
t := []string{"g", "h", "i"}
fmt.Println("dcl:", t)
```

```go
var fruits = []string{"apple", "grape", "banana", "melon"}
fmt.Println(fruits[0])
```

- One is the builtin `append`, which returns a slice containing one or more new values. Note that we need to accept a return value from `append` as we may get a new slice value.
```go
s = append(s, "d")
s = append(s, "e", "f")
fmt.Println("apd:", s)
```
- Slices can also be `copy`’d. Here we create an empty slice `c` of the same length as `s` and copy into `c` from `s`.
```go
c := make([]string, len(s))
copy(c, s)
fmt.Println("cpy:", c)
```
- "cutting" the data
```go
l := s[2:5]
fmt.Println("sl1:", l)

l = s[:5]
fmt.Println("sl2:", l)

l = s[2:]
fmt.Println("sl3:", l)
```

**Map**
- _Maps_ are Go’s built-in [associative data type](https://en.wikipedia.org/wiki/Associative_array) (sometimes called _hashes_ or _dicts_ in other languages).
- Map adalah tipe data asosiatif yang ada di Go, berbentuk key-value pair. Untuk setiap data (atau value) yang disimpan, disiapkan juga key-nya. Key harus unik, karena digunakan sebagai penanda (atau identifier) untuk pengaksesan value yang bersangkutan.
```go
var m map[string]int
m = map[string]int{}
fmt.Println("map:", m)
```

```go
m := make(map[string]int)
fmt.Println("map:", m)
```

```go
n := map[string]int{"foo": 1, "bar": 2}
fmt.Println("map:", n)
```

- assign new data
```go
m["k1"] = 7
m["k2"] = 13

fmt.Println("map:", m)
```
- declaring variable for map value
```go
v1 := m["k1"]
fmt.Println("v1: ", v1)
```
- The builtin `len` returns the number of key/value pairs when called on a map.
```go
fmt.Println("len:", len(m))
```
- The builtin `delete` removes key/value pairs from a map.
```go
delete(m, "k2")
fmt.Println("map:", m)
```


**Struct**
- Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu. Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti map, hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda.
- Dari sebuah struct, kita bisa buat variabel baru, yang memiliki atribut sesuai skema struct tersebut. Kita sepakati dalam buku ini, variabel tersebut dipanggil dengan istilah object atau object struct.

- Go’s _structs_ are typed collections of fields. They’re useful for grouping data together to form records.

```go
type person struct {
	name string
	age int
}

func main() {
	var p1 person
	p1.name = "John"
	p1.age = 42

	fmt.Println("Name:", p1.name)
	fmt.Println("Age:", p1.age)
}
```

- embedded struct
```go
type person struct {
	name string
	age int
	job
}

type job struct {
	company string
	roles string
}

func main() {
	var p2 person
	p2.name = "Doe"
	p2.age = 30
	p2.company = "Madeup Company"
	p2.roles = "Software Engineer"

	fmt.Println("Name:", p2.name)
	fmt.Println("Age:", p2.age)
	fmt.Println("Company:", p2.company)
	fmt.Println("Job Roles:", p2.roles)
}
```

- combine struct and slice
```go
type person struct {
	name string
	age int
}

func main() {

	var allEmployee = []person{
		{name: "Ethan", age: 54},
		{name: "Anderson", age: 63}
	}

	for _, employee := range allEmployee {
		fmt.Println(employee.name, " age is ", employee.age)
	}
}
```

## Defer
- Defer digunakan untuk mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("John")
	fmt.Println("Who are you?")
}
```


---
## References
1. https://github.com/golang/go
2. https://www.binaracademy.com/blog/apa-itu-golang-dan-fungsinya
3. https://www.freecodecamp.org/news/what-is-go-programming-language/
4. https://stackoverflow.com/questions/47347191/go-error-unexpected-nul-in-input
5. https://gobyexample.com/
6. https://dasarpemrogramangolang.novalagung.com/