// the places in code where a variable can be accessed
// a variable is first checked first inside the most inner block; if it's not found it's checking the next block above
package main

import "fmt"

var x = 4

func f() {
	fmt.Printf("%d", x)
}

func g() {
	var x = 5
	fmt.Printf("%d", x)
}

func main() {
	x = 6
	f()
	g()
	fmt.Printf("%d", x)
}
