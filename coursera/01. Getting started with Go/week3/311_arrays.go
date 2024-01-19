package main

import "fmt"

/*
Arrays
- Fixed length series of elements of a chosen type
- Elements accesed using subscript notation []
- Indices start at 0
- Elements initialized to zero value

Array Literal
- an array predefined with values
  var x [5]int = [5]{1,2,3,4,5}
- it can also infer size by using [...]
  x := [...]int{1,2,3,4}

Iterating through arrays
- use a for loop with the range keyword
- range returns two values: index and element at index


*/

var x [5]int

func main() {

	x[0] = 2

	y := [...]int{1, 2, 3, 4}
	z := []string{"a", "b", "c"}

	fmt.Println(x[0], x[1], y, z[2])

	for i, v := range y {
		fmt.Printf("ind %d, val %d\n", i, v)
	}
}
