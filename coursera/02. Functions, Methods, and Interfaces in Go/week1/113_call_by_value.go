package main

import "fmt"

/*
Call by value (default in Go)
- Passed arguments are copied to parameters; This means that the function cannot interfere with the original values

Tradeoffs of Call by value
- Advantage -> data encapsulation: function variables only changed inside the function
- Disadvantage -> copying time: large objects may take a long time to copy

Call by reference (alternative, have to do it manually)
- Programmer can pass a pointer as an argument
- Called function has direct access to caller variable in memory

Tradeoffs of Call by reference
- Advantage -> copying time
- Disadvatange -> data encapsulation: a bug inside the function can alter the variable inside main

*/

// Call by value
func foo(y int) {
	y = y + 1
	fmt.Println(y)
	// return y
}

// Call by reference
func foo2(y *int) { // the function takes a pointer to an integer as an argument
	*y = *y + 1
	fmt.Println(*y)
}

func main() {

	x := 2

	foo(x)
	fmt.Println(x)

	foo2(&x) // passes a pointer of x, the location in memory where is x
	fmt.Println(x)

}
