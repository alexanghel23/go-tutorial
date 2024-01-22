package main

import "fmt"

/*
Interfaces
- concept used in go that helps us getting polymorphism
- we don't need inheritance, overriding - we can use interfaces to accomplish the same thing

Interfaces are a set of method signatures
- name, parameters, return values
- implementation is NOT defined

Interfaces are used to express conceptual similarity between types
Example: Shape2D interface; All 2D shapes must have Area() and Perimeter()

Satisfying an Interface
- type satisfies an interface if type defines all methods specified in the interface
- Rectangle and Triangle types satisfies the Shape2D interface because they have Area() and Perimeter()
- additional methods are ok

- Similar to inheritance with overriding

*/

type Shape2D interface {
	Area() float64
	Perimeter() float64
}

type Triangle int

func (t Triangle) Area()      { fmt.Println("B") }
func (t Triangle) Perimeter() { fmt.Println("V") }
