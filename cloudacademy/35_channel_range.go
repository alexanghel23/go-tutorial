package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	squares := make(chan int, 20)

	for i := f(); i <= 100; i = f() {
		squares <- i
	}

	close(squares)

	for elem := range squares {
		fmt.Println(elem)
	}

}

/*
Channels can be iterated over. That is, you can use the range keyword in the same manner as you would when using it with arrays, slices, and/or maps.
This allows you to quickly and easily iterate over the messages that exist within a channel.
In the example provided here, on lines five to 11, I've created a squares function that returns another function which implements returning a sequence of squares.

In the main function on lines 17, 18 and 19, I use a simple for loop to load up the squares channel with all the squares of whole numbers where the square
is less than or equal to 100. Finally, on lines 23 to 25 I use another for loop with the range keyword to iterate over the current contents of the squares channel.
Note, it's very important to close a channel before iterating over the messages within the channel.
Without doing this, the program would eventually panic due to a deadlock.

Closing the channel in this example also highlights an important point about closing channels, and that is you can still read from them after
they are closed, you just can't continue to send messages to them if they have been closed. Running this example results in the following output.
Let's now comment out the close function on line 21 and then rerun the program. This results in a deadlock, as earlier explained.

In summary, you've just observed how to use the range keyword with a for loop to iterate over the messages within a named channel,
how closing a channel before iterating over it is important to avoid deadlocking, and that you can still read from a channel after it has been closed.
*/
