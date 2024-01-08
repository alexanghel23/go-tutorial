package main

import (
	"fmt"
)

var pl = fmt.Println

// recursion - when a function calls itself
// must have a condition that ends this calling

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

func main() {

	pl("Factorial 4 =", factorial(4))

}
