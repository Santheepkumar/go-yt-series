package main

import "fmt"

func main() {
	fmt.Println("=========== Basic Declaration ===========")
	var name string = "Software School"
	name = "Software School Tamil"
	var age int = 26
	var isCool bool = false

	fmt.Println(name, age, isCool)

	fmt.Println("=========== Short Declaration ===========")
	// var count int = 100
	count := false
	fmt.Println(count)

	fmt.Println("=========== Zero Values ===========")
	var num int        // 0
	var text string    // ""
	var isAwesome bool // false
	fmt.Println(num, text, isAwesome)

	fmt.Println("=========== Multiple Declaration ===========")
	var x, y, z string = "A", "B", "C"
	fmt.Println(x, y, z)

	fmt.Println("=========== Constants ===========")
	const a int = 12
	fmt.Println(a)

	fmt.Println("=========== Type conversion ===========")
	const number int = 100
	numberFloat := float32(number)
	numberFloat = 100.1111

	fmt.Println("numberFloat", numberFloat)

	fmt.Println("=========== Other Derived Types ===========")

	// Array
	// Slice
	// Maps
	// Structs
}
