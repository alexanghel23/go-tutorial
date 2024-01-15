package main

import "fmt"

type company struct {
	name string
}

func main() {
	var x interface{} = company{"CloudAcademy"}

	c1 := x.(company)
	fmt.Println(c1)

	if c2, ok := x.(company); ok {
		fmt.Println(c2, ok)
	}

	//n := x.(int); n++ //runtime panic
}

/*
Go provides you with the ability to perform type assertions on interface types. A type assertion provides access to an interface's underlying concrete type.

The type assertion "x. " asserts that the concrete value stored in x is of type T, and that x is not nil. In the example provided here,
a custom company struct type is declared on lines 5 to 7. Then, within the main function on line 10, the variable x is declared using the empty interface
type and is then assigned an instance of the company struct type.

On line 12, short notation is used to cast x to its concrete type and assign the cast to the variable c1, which is then immediately printed out on line 13.
Lines 15 to 17 demonstrate how to do the exact same cast but in a way that allows the application to keep running should the cast fail.
Here the type assertion is checked to see if the assertion was true or false. If you don't use this technique to test your assumed type assertion and it
turns out to be incorrect, a runtime panic is thrown and the program could terminate if the panic isn't recovered from, as is the case on line 19.
Let's run this example as is, and as expected the program completes without any panics being encountered.

Next let's uncomment line 19. Now this time this should result in a panic, since this particular cast will not work. Rerunning the program indeed results in the expected crash.

In summary you've observed how to implement type assertions, how to test type assertions, allowing the application to have alternate path flows depending on
whether the type assertion was correct or not, and how unhandled and incorrect type assertions cause runtime panics resulting in possible program termination.

*/
