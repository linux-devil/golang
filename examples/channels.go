package main

import (
	"fmt"
)

func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	buffered_messages := make(chan string, 2)

	buffered_messages <- "ping1"
	buffered_messages <- "ping2"

	fmt.Println(<-buffered_messages)
	fmt.Println(<-buffered_messages)

}
