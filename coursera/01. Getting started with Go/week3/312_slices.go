package main

import "fmt"

/*
Slices
- A window on an underlying array
- variable size, up to the whole array
- Three properties:
  - Pointer: indicates the start of the slice
  - Length:  number of elements in the slice
  - Capacity: max number of elements
- Functions:
  - len()
  - cap()

Accesing slices
- Writing to a slice changes underlying array
- Overlapping slices refer to the same array elements

Slice literals
- can be used to initialize a slice
- creates the underlying array and creates a slice to reference it
- slice points to the start of the array, length is capacity


*/

func main() {

	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}

	s1 := arr[1:3]
	s2 := arr[2:5]

	fmt.Println(s1, s2)

	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))

	// Accesing slices
	fmt.Println(s1[1], s2[0])

	// Slice literals - there is nothing in the brackets and so the compiler knows this is a slice. Then it creates the underlying array and next it makes the slice point to the entire array
	sli := []int{1, 2, 3}
	fmt.Println(sli)

}
