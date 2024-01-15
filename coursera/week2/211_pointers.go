// Pointer: address to a data in memory
// & operator return the address of a variable/function
// * operator returns data at an address (dereferencing) when added before of a pointer (address)

// new() - function that creates a variable and returns a pointer to the variable; variable is initially initialized to zero

package main

import "fmt"

var x int = 1
var y int

var ip *int // ip  is not actual integer, it is a pointer to integer;

func main() {

	ip = &x

	fmt.Println(x, ip, y)

	y = *ip

	fmt.Println(x, ip, y)

	ptr := new(int)
	*ptr = 3

	fmt.Println(ptr, *ptr)
}
