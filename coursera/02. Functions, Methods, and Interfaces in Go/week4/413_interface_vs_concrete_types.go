package main

import "fmt"

/*
Concrete types
- specify the exact representation of the data and methods
- complete method implementation is include

Interface Types
- specifies some method signatures
- implementations are abstracted

Interface values
- can be treated like other values (assigned to variables, passed, returned)
- they have two components:
	- Dynamic type: concrete type which it is assigned to
	- Dynamic value: value of the dynamic type

*/

type Speaker interface {
	Speak()
}

type Dog struct {
	name string
}

func (d Dog) Speak() {
	fmt.Println(d.name)
}

func main() {
	var s1 Speaker
	var d1 = Dog{"Brian"}
	s1 = d1
	s1.Speak()

}
