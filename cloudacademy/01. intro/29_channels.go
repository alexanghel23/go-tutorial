package main

import "fmt"

func main() {
	msgChan := make(chan string)

	go func() {
		msgChan <- "Cloud"
		msgChan <- "Academy"
		msgChan <- "2020"
	}()

	msg1 := <-msgChan
	msg2 := <-msgChan
	msg3 := <-msgChan

	fmt.Println(msg1, msg2, msg3)
}

/*
Channels are communication objects used by goroutines to communicate and share data with each other. In their simplest incarnation, one goroutine
will write messages into the channel, while another goroutine will read the same messages out of the channel. Channels are created using the make
function together with the chan keyword. Like variables, channels are typed and as such can be used to transport messages only of that data type, for which was declared with.

To write data into a channel, you first specify the name of the channel followed by the left arrow syntax, and then the data itself.
This style of syntax is highly readable, since the data always flows in the direction of the arrow, making the concept of message passing and processing
quite easy to understand.

Reading data off the channel involves moving the left arrow syntax to the front of the channel name. In the example shown here, the main function first
declares a message channel of type string on line six. Next, an anonymous goroutine is declared on lines spanning eight through to 12.
Inside this anonymous goroutine, three individual messages are written into the channel.
The main function then concludes by reading back off the channel, each of the three messages.

Let's run this example and examine the resulting output. And, as expected, we get the result Cloud Academy 2020.

In summary - you've observed; How to use the make function together with the chan keyword to create typed channels.
How to write messages onto the channel. And then how to read the messages back off the channel. Note: Channels act as first-in-first-out queues.
For example, if one goroutine writes messages into a channel and a separate goroutine reads them off the channel, then the messages will be
received in the exact same order sent.
*/
