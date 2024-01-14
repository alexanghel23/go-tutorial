package main

import (
	"fmt"
	"time"
)

func in(channel chan<- string, msg string) {
	channel <- msg
}

func out(channel <-chan string) {
	for {
		fmt.Println(<-channel)
	}
}

func main() {
	channel := make(chan string, 1)

	go out(channel)

	for i := 0; i < 10; i++ {
		in(channel, fmt.Sprintf("cloudacademy - %d", i))
	}

	time.Sleep(time.Second * 10) //crude
}

/*
When using channels as function parameters, as you often will, by default can send and receive within the function.

To provide additional safety at compile time, channel function parameters can be defined with a direction. That is, they can be defined to be read-only or write-only.
The example provided here shows you how to accomplish this. The in function signature defined on line eight has a channel function parameter
locked down for writing only, as per the additional channel direction notation inserted between the channel name and the channel type.

Likewise, the out function signature declared on line 12 has a channel function parameter locked down for reading only.

In summary, you have observed how to use channel direction notation to control how channels are used inside a function.
*/
