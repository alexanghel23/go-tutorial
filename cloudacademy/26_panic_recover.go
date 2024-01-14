package main

import "fmt"

func system() int {
	fmt.Println("system started...")

	defer func(msg string) {
		if r := recover(); r != nil {
			fmt.Println("recovered")
		}
		fmt.Println(msg)
	}("blah")

	var data []int
	var x = data[0] //causes runtime spanic!
	x++

	fmt.Println("system finished!")

	return 1
}

func main() {
	data := system()
	fmt.Println(data)

	panic("die!") //exits program with non-zero code
}

/*
Go provides additional capabilities when exceptions occur in the form of panics and recovery.
Custom panics can be thrown when something exceptional and unexpected has happened.
Panics, if not recovered from, will cause the program's execution to be exited with a non-zero exit code, which the operating
system uses to determine that something abnormal has just happened.
A panic can be recovered from if and only if, the surrounding function in which the panic happened has a deferred function that itself
then calls the inbuilt recover function and then does something appropriately to resume.

In the example provided here, the system function will throw a runtime panic on Line 16. Since we're trying to access the data slice, we know it's still empty.
The anonymously declared defer function on Lines 8 to 13 will now take program execution from the panic, and since it calls the recover
function within itself, it enables the panic to be recovered from.
Interestingly, because the system function return statement didn't get hurt, the returned value is received back within the main
function on Line 25 will be zero and not one. Zero being the zero value for the data type, int, which is the date type for the returned value.

Finally, on Line 28 for demo purposes only, I manually invoked panic with the message, die. And since there is no recovery option available,
the program execution exits with a non-zero exit code. Let's now run this example and examine the resulting output.
Here, we can see the following messages are printed out: system started, recovered, blah, and zero.
Before, the program then finally halts due to the panic, die, exiting with the status code of 2.

Therefore, in summary, you've just observed how to handle and recover from runtime panics, using the inbuilt recover function,
how to create and raise panics, and how an unhandled panic will cause program termination with a non-zero exit code.
*/
