package main

import "fmt"

/*
Struct
- aggregate data type
- groups together other objects of arbitrary type

Initializing structs
- can use new()
- initializes fields to zero
	p1 := new(Person)
- can initialize using a struct literal
	p1 := Person{name: "joe", ...}

*/

type Person struct {
	name  string
	addr  string
	phone string
}

var p1 Person

func main() {

	p1 = Person{"Alex", "Tarlesti", "0723936090"}
	fmt.Print(p1)

	fmt.Println(p1.name)

	p2 := new(Person)
	p2.name = "Gogu"
	p2.addr = "Posesti"
	p2.phone = "078884452"

	fmt.Println(p2)

	p3 := Person{name: "Gina", addr: "asa", phone: "123"}
	fmt.Println(p3)
}
