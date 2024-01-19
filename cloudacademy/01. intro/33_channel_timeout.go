package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go func(channel chan string) {
		time.Sleep(5 * time.Second)
		channel <- "cloudacademy"

	}(channel)

	select {
	case msg1 := <-channel:
		fmt.Println(msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout!")
	}
}

/*
Timeouts can be specified in select statements to help manage situations where it's taking too long to receive a message from any one of the channels being monitored.

Consider using timeouts when you're implementing something that connects to an external resource. When working with external resources, you can never
guarantee the response times and, therefore, may need to proactively take action after a predetermined timeout. Implementing a timeout with a select
case statement is very straightforward. Simply import the time package and then have a case statement that unblocks on the time.after channel.
With this configuration in place, the select statement will timeout and unblock after a configured amount of time if no other channels liven up.

Let's see this in action. Running this program results in the message Timeout. The explanation for this follows. We start up an anonymous
go routine on lines nine through to 13.

Within it, we immediately pause for five seconds. Regardless, the main program flow continues onto the select case statements implemented across lines 15 to 20.
The last case statement declares a two-second timeout using the time.after function. And the result follows that the message Timeout is returned
in the program's output, as seen here. Let's swap the timings on lines 10 and 18 and then rerun the program.
This time, the timeout is not exceeded, and the program instead prints out the message cloudacademy.

In summary, you've just observed how a timeout can be declared within a select case statement and how to utilize the time.after function to implement
a timeout when working with channels and messages.
*/
