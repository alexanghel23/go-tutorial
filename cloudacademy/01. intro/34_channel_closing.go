package main

import (
	"bytes"
	"fmt"
)

func process(work <-chan string, fin chan<- string) {
	var b bytes.Buffer
	for {
		if msg, notClosed := <-work; notClosed {
			fmt.Printf("%s received...\n", msg)
			b.WriteString(msg)
		} else {
			fmt.Println("channel closed")
			fin <- b.String()
			return
		}
	}
}

func main() {
	work := make(chan string, 3)
	fin := make(chan string)

	go process(work, fin)

	word := "cloudacademy"

	for j := 0; j < len(word); j++ {
		letter := string(word[j])
		work <- letter
		fmt.Printf("%s sent...\n", letter)
	}

	close(work)

	fmt.Printf("result: %s", <-fin)
}

/*
Go channels can be explicitly closed to help with synchronization issues. Closing a channel is done by invoking the built-in close function,
providing it with a channels name that you want to close.

Let's quickly run this example, and then next explain its inner workings. After running it, we can see that it has printed out the following sequence of messages.
Within the example displayed here, the process function is started as a go routine on line 26. The process function takes in two channels as inputs, work and fin.
The work channel is used to receive messages as work to perform.

The fin channel is used to communicate that all required work has been completed or finished. The work channel receives messages that are individual characters
extracted one at a time from the word string variable initialized in the main function on line 28. Within the process function, a non-terminating for loop is started.
Each cycle of the loop waits until a new message arrives on the work channel. In this case, the process function is designed simply to restitch all of the
individual characters received back into a single string.
This is then returned when the channel is determined to have been closed, which is done on line 16, with the final result being
printed out by the Printf function on line 38. This correlates with the result seen in the output at the bottom, here.

In summary, you've observed how to use the close function to close a channel, which is useful for synchronization purposes.
*/
