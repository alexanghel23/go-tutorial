package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("cloud")

	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("academy")

	wg.Wait()
}

/*
Goroutines can also easily be declared using anonymously declared functions, as is demonstrated here.
One off blocks of code can be launched as a goroutine to ensure that they are non-blocking against the remainder of the program.
This can be seen on lines 12 to 15 and then again on lines 18 to 21. Additionally, in this example I've introduced the concept of a WaitGroup.

A WaitGroup is used to help the overall program manage and halt itself until all tagged goroutines have returned.
Setting up and implementing a WaitGroup is a fairly straightforward process. First of all import the sync package as seen on line five.
Next declare a WaitGroup variable and then call the Add function once per launched goroutine, seen on lines 11 and 17.

Within each goroutine, defer a call to the Done function on the WaitGroup variable as per lines 13 and 19.
Finally, in the main function thread, which in this case is the end of the main function, make a call to the Wait function on the WaitGroup variable.
Launching this example, we can examine the resulting output. With everything in place, the WaitGroup manages the goroutines and knows exactly how long
to halt the main thread to allow all tagged goroutines enough time to complete their execution, as can be seen by the produced output.

In summary, you have observed how to create anonymously declared goroutines and how to use WaitGroups to halt the main program's execution until all goroutines have completed.
*/
