package main

import (
	"fmt"
)

var pl = fmt.Println

// structs allow you to store values with different data types

type rectangle struct {
	length, height float64
}

// function that is part of the struct
func (r rectangle) Area() float64 {
	return r.height * r.length
}

func main() {

	rect1 := rectangle{10.0, 15.0}
	pl("Rect area is", rect1.Area())
}
