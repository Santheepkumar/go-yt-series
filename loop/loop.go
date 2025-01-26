package main

import "fmt"

func main() {

	// For loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// While

	j := 1
	for j <= 5 {
		fmt.Println("While", j)
		j++
	}

	// Infinite loop

	k := 1
	for {
		fmt.Println("Infinite", k)
		k++
		break
	}

	// Range
	nums := []int{1012, 2, 3, 4, 5, 112}

	for index, value := range nums {
		fmt.Printf("Index: %d, Value %d\n", index+1, value)
	}
}
