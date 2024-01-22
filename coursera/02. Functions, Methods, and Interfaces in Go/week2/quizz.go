package main

import "fmt"

// quiz 1
func fA() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fB := fA()
	fmt.Print(fB())
	fmt.Print(fB())
}

/*
The result is not "11" because each time you call the closure `fB`, it maintains its own separate state.
Closures in Go capture the variables they reference, and in this case, the closure returned by `fA` captures the variable `i`.
Each time the closure is called, it increments and retains its own copy of `i`.

Let's break down the execution:

1. `fB := fA()` creates a closure with its own copy of the variable `i` (initialized to 0).
2. `fmt.Print(fB())` calls the closure, increments its copy of `i` to 1, and prints the result (1).
3. `fmt.Print(fB())` calls the closure again, increments its own copy of `i` (which is now 1) to 2, and prints the result (2).

If the closures were sharing the same state, the output would be "11" because both calls would be operating on the same `i`.
However, in Go, closures capture the variables they reference at the time they are created, creating independent copies for each closure instance.

In summary, closures in Go are functions that have access to variables from the surrounding lexical scope and can "close over" those variables, maintaining their own copies.
Each closure instance can have its own state, which is separate from other closures created from the same enclosing function.
*/

// quiz 2

// func main() {
// 	i := 1
// 	fmt.Print(i)
// 	i++
// 	defer fmt.Print(i + 1)
// 	fmt.Print(i)
// }
