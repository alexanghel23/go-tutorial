package main

import "fmt"

func extend(val string) func() string {
	i := 0
	return func() string {
		i++
		return val[:i]
	}
}

func main() {
	ca := "cloudacademy"

	word := extend(ca)

	for i := 0; i < len(ca); i++ {
		fmt.Println(word())
	}
}

/*
Functions in Go are considered first-class citizens of the language since they can be declared and held in variables and can even be declared as the return parameter of another function.
In the example shown here, the extend function on line five is declared to return a function itself.
The return to function is anonymously defined within the body of the extend function on lines seven through to 10.
Within the main function on line 16, the extend function is called with the ca string variable which contains the string cloudacademy.
The extend function returns a function which is assigned to the variable word.

A for loop is then used to repeatedly call and invoke the function currently stored in the word variable.
The outcome of running this application results in a list of printed statements starting with just the character c, followed by the character sequence c-l on a new line, then followed by the character sequence c-l-o on a new line, and so on until the full string cloudacademy is printed out.

The anonymous function returned by the extend function is an example of a closure.
A closure is a special type of anonymous function that has the ability to reference variables declared outside of itself.
In this case, the inner returned function is referencing the variable i designed to be a counter.
Now, this is quite similar to how functions can access and potentially manipulate global variables which is strongly argued is not good practice.

To address this, closures have one other defining ability and that is they can continue to reference variables that they had access to during creation even if those variables are no longer referenced elsewhere.
In other words, a closure gives you access to an outer function's scope from within an inner function.
This again is demonstrated by the variable i, declared on line six within the extend function.
It holds state and is referenced and incremented within that inner returned function.

In summary, you have observed the following: that functions are first-class citizens within the Go programming language, how to assign anonymous functions to variables, how to create functions that return other functions, and how to use anonymous functions to create closures.
*/
