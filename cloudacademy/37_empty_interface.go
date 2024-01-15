package main

import "fmt"

type company struct {
	name string
}

func main() {
	var a, b, c, d interface{}

	a = 42
	b = "blah"
	c = true
	d = company{"cloudadacademy"}

	func(list ...interface{}) {
		for _, v := range list {
			fmt.Printf("%v, %T\n", v, v)
		}

	}(a, b, c, d)
}

/*
The empty interface type in Go denoted by the keyword interface and empty curly brackets is used to indicate
that a variable can hold values of any type, since by definition every type implements at least zero methods.
The empty interface is often useful when you need to accommodate an unknown set of types that a function may consume.
As per the example given here, the anonymous variadic function declared on lines 17 to 22 takes in a list typed as the empty interface.

Internally to the function, the list is then iterated across using the range keyword, with each value and its concrete type printed
out using the Printf function on line 19. Running this example results in the following output.
Being able to call a function with as many variables, of differing types, helps to reduce code bloat by condensing and generalizing many
functions into a singular multipurpose function.
Compiler type safety checking is given up when taking this approach, so you need to consider your requirements carefully when doing this.

In summary, you've observed: how to use and declare variables that are of type empty interface and that all Go types are of type empty
interface since by definition every type implements at least zero methods.

*/
