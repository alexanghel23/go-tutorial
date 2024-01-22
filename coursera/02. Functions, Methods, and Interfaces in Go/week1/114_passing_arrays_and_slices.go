package main

import "fmt"

/*
Passing array arguments
- array arguments are copied
- arrays can be big, so this can be a problem

Passing array pointers
- it's possible to pass array pointers
- passing an array copies the entire array

Pass slices instead
- in general, you should use slices in go instead of arrays
- slices contain a pointer to an array
- passing a slice copies the pointer
- a slice is basically a structure that contains 3 things: a pointer to an array, length and capacity.
- when you pass a slice you copy directly the pointer, so we don't have to reference it and dereference it like we did with arrays

*/

func foo(x [3]int) int {
	return x[0]
}

func foo2(x *[3]int) {
	(*x)[0] = (*x)[0] + 1
	// return x[0]
}

func fooslice(sli []int) {
	sli[0] = sli[0] + 1
}

func main() {

	// array
	a := [3]int{1, 2, 3}
	fmt.Println(foo(a)) // here the function has to copy the array into x (the parameter in foo)

	b := [3]int{1, 2, 3}
	foo2(&b)
	fmt.Println(b)

	// slice
	c := []int{1, 2, 3}
	fooslice(c)
	fmt.Println(c)

}
