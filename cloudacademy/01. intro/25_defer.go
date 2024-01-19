package main

import "fmt"

func doSomething(msg string) {
	fmt.Println(msg)
}

func system() int {
	fmt.Println("system started...")

	defer doSomething("cleanup")
	defer doSomething("stop")

	fmt.Println("system finished!")

	return 1
}

func main() {
	data := system()
	fmt.Println(data)
}

/*
The defer statement is used to defer the execution of a function until the currently running and surrounding function's execution returns.

A motivation for which you may consider using defer could be when you're working with or processing external resources.
For example, files on the file system, or perhaps database connections. For example, you may have a requirement to open a file and manipulate it.
By calling defer immediately after opening it, you can be assured that it gets closed and written back to the file system. once the surrounding
function or method completes, regardless of the execution path taken within it.
Calling defer also has the added benefit of insuring that you don't unintentionally leak resources due to unforeseen program flow,
since your resource management is explicitly declared, right from the start.

Using a defer statement ensures that the deferred function is always called, regardless of the execution path within the surrounding function.
It's also possible to stack up multiple defer statements and then have them processed in LIFO order, last in, first out.
Again, only after the surrounding function returns.
Within the example system function seen here, defer is used on lines 12 and 13 to ensure that the do something function is executed
immediately after the system function execution completes, with the messages "clean up" and "stop" respectively.

Now, when this program is executed, the order of the printed items will be "system started", "system finished", "stop", "clean up," and "one."

In summary, you have just observed that deferred functions are called after the surrounding function returns, how multiple
deferred functions are called in LIFO order, and that deferred functions are useful for cleaning up resources.
*/
