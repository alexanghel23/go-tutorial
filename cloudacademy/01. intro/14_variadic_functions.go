package main

import "fmt"

func displayCount(id int, letters ...string) {
	count := 0
	for range letters {
		count++
	}

	//display id, letters count, letters type, and letters content
	fmt.Printf("%d - %d - %T - %s\n", id, count, letters, letters)
}

func main() {
	displayCount(1, "c", "l", "o", "u", "d")
	displayCount(2, "a", "c", "a", "d", "e", "m", "y")

	cloud := []string{"d", "e", "v", "o", "p", "s"}
	displayCount(3, cloud...)
}

/*
Functions can also be defined with a variable length trailing or last input parameter. These types of functions are referred to as Variadic Functions.

It's important to understand that a variadic input parameter must be declared as the last parameter in the function signature.
A variadic input parameter will be passed into the function as a slice. Let's see how this works.
The displayCount function implemented on line five is declared with two parameters.
Notice here that the string data type is preceded by three ellipses. This informs the compiler that the letters input parameter is a variadic input param, it is variable in length.
Back within the main function, lines 16 and 17 show you how to call the displayCount function with a variable length set of input parameters.

The very first parameter is passed to the id input parameter with all remainder strings being wrapped up and passed to the letters variadic input param which becomes a slice of string.
The letters slice is iterated over using the range keyword. You can see this on lines seven, eight, and nine. In this case I'm simply implementing a basic count for demonstration purposes.

Another interesting Go programming technique to consider when calling variadic functions is implemented on lines 19 and 20.
Line 19 declares a slice of strings which is then passed into the displayCount function on line 20, where it is first unpacked using the 3 ellipses with the result being the displayCount function being executed in the same manner as previously described.

In summary, you have observed how to declare a variadic function which takes a variable list of input arguments for the last parameter, how to call variadic functions including using a slice and unpacking it and how within a variadic function, the last input parameters type is a slice and how you can use the range keyword to iterate over its contents.
*/
