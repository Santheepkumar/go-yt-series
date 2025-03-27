package main

import "fmt"

func changeIt(p *int) {
	*p = 42 // Modifies the original variable
}

func main() {
	x := 10
	changeIt(&x)
	fmt.Println(x) // Output: 42
}
