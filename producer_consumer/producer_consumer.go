package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("Producing %d \n", i)
		ch <- i // message passing or producing
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch chan int) {
	for item := range ch { // consuming hero
		fmt.Printf("Consuming %d \n", item)
	}
}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	go producer(ch, 5)
	consumer(ch)

	go producer(ch2, 10)
	consumer(ch2)
}
