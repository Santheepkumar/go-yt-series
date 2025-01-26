package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello %s", name)
}

func add(x, y int) int {
	return x + y
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return a / b, nil
}

func arbitrarySum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {

	// Basic Functions
	// greet("Bruce")

	// Parameter function
	// sum := add(10, 20)
	// fmt.Println(sum)

	// Functions with multiple return values
	quotient, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", quotient)
	}

	// Arbitrary functions or variadic functions
	sum := arbitrarySum(1, 2, 3)
	fmt.Println(sum)

	// Anonymous functions
	greet := func(name string) {
		fmt.Printf("Hello %s", name)
	}
	greet("You \n")

	// Closures

	// first function declaration, return type of declared function, return type of returned function
	adder := func(x int) func(int) int {
		return func(y int) int {
			return x + y
		}
	}

	multiplexer := func(x int) func(int) int {
		return func(y int) int {
			return x * y
		}
	}

	addFive := adder(5)
	fmt.Println("addFive", addFive(10))

	multiplyFive := multiplexer(5)
	fmt.Println("multiplyFive", multiplyFive(10))

	// Named return

	result := multiply(100, 2)
	fmt.Println(result)
}
