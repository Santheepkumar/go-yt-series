package main

import "fmt"

func main() {

	age := 17
	if age >= 80 {
		fmt.Println("You're a major")
	} else if age >= 18 {
		fmt.Println("You're an adult")
	} else {
		fmt.Println("You're a minor")
	}

	// Inline initialization

	if num := 9; num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}
