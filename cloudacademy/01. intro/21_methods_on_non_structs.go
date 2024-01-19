package main

import (
	"fmt"
	"strings"
)

type upstring string

func (msg upstring) up() string {
	s := string(msg)
	return strings.ToUpper(s)
}

func main() {
	message := upstring("cloudacademy")
	fmt.Println(message.up())
}

/*
The coding example displayed here demonstrates how to implement a method with a non-struct typed receiver argument.

In this example, the up method defined on lines eight to 11 implements a simple string upper casing function.
The receiver type is of type upstring, and is a custom non-struct string based type declared on line six.
With this in place, the up method can now be invoked using dot notation.

This is demonstrated within the main function on line 15.
Let's now run this example and see what happens.
And, as expected, the output is the string cloudacademy uppercased. Changing the cloudacademy message to devops and rerunning the program again provides the expected result.

In summary, you've just observed how receiver arguments can be declared as non-struct types.
*/
