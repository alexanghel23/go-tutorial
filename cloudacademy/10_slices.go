package main

import "fmt"

func main() {
	arr1 := [6]int{1, 3, 5, 7, 11, 13}
	slice1 := []int{1, 3, 5, 7, 11, 13}
	slice2 := slice1[1:2]
	var slice3 = make([]int, 2, 3)

	fmt.Println(arr1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	slice3 = slice1[1:4]

	fmt.Println(slice3)
	fmt.Println(len(slice3))

	slice1 = append(slice1, 200, 300, 400)
	fmt.Println(slice1)

	slice2 = append(slice2, []int{7, 8, 9}...)
	fmt.Println(slice2)

	copyslice := make([]int, len(slice1))
	copy(copyslice, slice1)
	fmt.Println(copyslice)
}

/*
Slices build on the capabilities of arrays with their point of difference being that they are resizable which means you can expand and contract them at runtime.
In fact because of this additional flexibility, you will more likely see slices being used and implemented than arrays.
Declaring a slice can be performed one of three ways.

The first way is to declare a slice in very much the same way as you would an array with the exception being omitting the size as can be seen on line seven.

The second way to declare a slice is to do so by taking a slice of another array or slice as shown on line eight.

The third and final way to declare a slice is to do so by using the make function. The make function takes three input parameters, the type of the slice, the length of the slice, and the capacity of the slice.

This is demonstrated on line nine. Line 16 resizes slice3 by reassigning it to a slice of slice1.
Line 21 demonstrates how to use the append function to add additional elements to slice1.
Line 24 demonstrates how to do the same thing, but where a slice literal is passed in as the second argument to the append function and instructed to be unpacked using the three ellipses.
Lines 27 and 28 show how to use the copy function to copy one slice into another.

In summary, you've just observed that slices build on the capabilities of arrays with their main point of difference being that they are resizable, that it is more likely to see slices in Go source code than it is to see arrays, and that there are three different ways in which slices can be created.

*/
