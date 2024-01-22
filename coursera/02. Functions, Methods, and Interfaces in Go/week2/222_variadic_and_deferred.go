package main

import "fmt"

/*
Variadic functions
- Functions can take a variable number of arguments
- Use ellipsis ... to specify
- Treated as a slice inside function
- Can also pass a slice to variadic function (need the ...suffix)

Deferred function calls
- call can be deferred until the surrounding function completes
- typically used for cleanup activities
- arguments of a deferred call are evaluated immediatialy and not later when the call actually happens

*/

func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main() {

	fmt.Println(getMax(1, 23, 55, 0, -2, 152, 7))

	vSlice := []int{1, 5, 9, 15, 4, 2, 47, 21, 33}

	fmt.Println(getMax(vSlice...))

	defer fmt.Println("Bye")
	fmt.Println("Hello")

	i := 2
	defer fmt.Println(i + 1)
	i = i + 5
	fmt.Println("Hey")
}
