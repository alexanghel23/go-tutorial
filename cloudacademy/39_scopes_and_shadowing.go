-- util.go --
package main

var x int = 11

-- main.go --
package main

import "fmt"

func print(id int, x int) {
	fmt.Printf("%d: x=%d\n", id, x)
}

func main() {
	print(1, x)
	
	x := 2

	print(2, x)

	func(x int){
		print(3, x)
		if x := 3; x == 3 {
			//x := 100
			print(4, x)
		}
		print(5, x)
	}(5)
	
	print(6, x)
}

/*
The following Go code emphasizes the concept of scope and variable shadowing. 

Lets first run this code sample so that you can correlate the explanation with the programs output. 
Here we can see the following output has been produced, within the main function there are six cores to the print function scattered 
throughout the implementation to print out the current value of the variable X, which undergoes shadowing due to the fact that the 
variable X is declared and initialized multiple times within different scopes. Within the util dot Go file variable X is first declared 
and initialized with the value eleven on line four. Note here that X is being declared in package scope within the main package which 
happens to span both the util dot Go, and the main dot Go files. 

Jumping down into the main function declared within the main dot Go file. We can see that X is re declared on line eighteen. 
This is permitted since we are doing it within a new scope, that being the main function scope. Lines twenty two to twenty nine define 
an anonymous function which takes a named input perimeter X, the anonymous function is called with the value five on line twenty nine. 
Therefore X takes on the value five within the anonymous function as it starts. Within this anonymous function an F statement is declared
 on line twenty four, which happens to declare X again it initializes it with the value three. 

This is an implicate blocked scoped declaration, and in this case the execution will enter into, the F statement since X is indeed three. 
Exiting back out of the F statement X retains the value five since the implicitly scoped X associated with the F statement is now deceases being no longer in scope. 
Finally as we exit out of the anonymous function X retains the value two, since the enormous function scope has been completed, and we're 
back within the main function scope. While we have this example loaded lets edit, end on comment line twenty five, within the F statement. 
This is an example of an explicitly scoped declaration of X within the F statement. Rewriting this example now changes the value printed 
out in the fourth print statement within the programs output. 

In summary you have observed how variables can be shadowed, and what happens when they are. How scope sits and controls the lifetime of a variable, 
how a same named variable can exist in different scopes. How packaged scopes can't be in multiple files, and how different 
scopes exist within a program, and can be nested.
*/