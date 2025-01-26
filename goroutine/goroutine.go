package main

import (
	"fmt"
	"time"
)

func Count() {
	i := 1
	for {
		fmt.Println(i)
		i++
	}
}

func main() {

	go Count()
	time.Sleep(5 * time.Second) // Blocked
	fmt.Println("Main function running.")

}
