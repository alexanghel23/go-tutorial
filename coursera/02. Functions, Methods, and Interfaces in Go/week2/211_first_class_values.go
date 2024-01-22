package main

import "fmt"

/*
Functions are First-class
- functions can be treated like any other type (variables can be declared with a function type)
- functions can be created dynamically
- functions can be passed as aruguments and returned as values
- functions can be stored in data structures

*/

// 1. Variable as functions
var funcVar func(int) int

// 2. functions as arguments
func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func incFn(x int) int {
	return x + 1
}

func decFn(x int) int {
	return x - 1
}

func main() {
	// 1. Variable as functions
	funcVar = incFn
	fmt.Println(funcVar(1))

	// 2. functions as arguments
	fmt.Println(applyIt(incFn, 2))
	fmt.Println(applyIt(decFn, 2))

	// 3. Anonymous functions

	v := applyIt(func(x int) int { return x + 1 }, 3) // it's basically the increment function, defined with no name
	fmt.Println(v)

}
