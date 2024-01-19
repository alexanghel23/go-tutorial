package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pause() {
	n := rand.Intn(3) // n will be between 0 and 3
	time.Sleep(time.Duration(n) * time.Second)
}

func doSomething(msg string) {
	pause()
	fmt.Println(msg)
}

func main() {
	rand.Seed(time.Now().Unix())

	doSomething("sync1")

	go doSomething("async1")
	go doSomething("async2")
	go doSomething("async3")

	doSomething("sync2")

	time.Sleep(time.Second * 10)
}

/*
Go provides several very cool language features that specifically address concurrency and asynchronous program flow requirements.

Fundamental to these features is the concept of a goroutine, which is implemented using the go keyword. Goroutines can be thought of as conceptually
like lightweight threads, although they are not strictly threads in the sense of what the operating system sees and manages.
Instead, goroutines are managed and scheduled by the Go runtime and done entirely within the virtual space of the Go runtime.
Goroutines enable you to start up and run other threads of execution concurrently within your program.
They are designed to be lightweight with minimal overhead, meaning you can literally spin up thousands of these without incurring performance penalties.

Let's take a quick look at an example that uses goroutines to async some of the program flow within it.
In the example shown here, the doSomething function declared on lines 14, 15, and 16 implements the sequence whereby it pauses itself for a
random amount of time, somewhere between zero and three seconds, and then resumes and prints out the inputted message.
Back within the main function, the doSomething function is called synchronously on lines 20 and 26, seeded with the messages sync1 and sync2, respectively.
On lines 22, 23, and 24, three goroutines are started, each with their own specific message, indicating that they are async executed.

Let's now run this example and examine the resulting output.
Here, we can see that the order of the messages is different to the order as seen in the code.
The reason for this is that the goroutines all run at approximately the same time concurrently. And since they are each required to pause for a random
amount of time, the order in which they wake up and resume will be different.
Note, the final time.Sleep function, as used on line 28, is a quick and basic approach to allowing the program's goroutines to complete
before the main program execution does. I'll show some better techniques as to how to manage this aspect when working with concurrency within Go as the course proceeds.

In summary, you have observed how to create goroutines using the go keyword to start concurrent flows of execution.
*/
