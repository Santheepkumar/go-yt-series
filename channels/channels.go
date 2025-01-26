package main

import "fmt"

func main() {
	messages := make(chan string)

	// go routine thread 2
	go func() {
		// Heavy task - Video compress
		messages <- "Hello from channel"
	}()

	// main go routine thread 1
	msg := <-messages

	fmt.Println("msg", msg)
}
