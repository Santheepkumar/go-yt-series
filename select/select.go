package main

import (
	"fmt"
	"time"
)

func sayHello(chName string) {
	fmt.Printf("Hello from %s", chName)
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Channel 1"
	}()
	go func() {
		ch2 <- "Channel 2"
	}()

	select {
	case msg1 := <-ch1:
		sayHello(msg1)
	case msg2 := <-ch2:
		sayHello(msg2)
	}

}
