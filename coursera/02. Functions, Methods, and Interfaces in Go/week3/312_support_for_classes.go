package main

import "fmt"

/*

- No "class" keyword in go

Associating methods with Data
- method has a receiver type that is associated with
- use dot notation to call the method
*/

type MyInt int

func (mi MyInt) Double() int { // this doen't work with built-in types
	return int(mi * 2)
}

func main() {
	v := MyInt(3)
	fmt.Println(v.Double()) // Object v is an implicit argument to the method. There is no explicit argument, but a hidden one. v object is passed to the function (call by value, so a copy of v is passed to double)
}
