package main

import "fmt"

type Person struct {
	name      string
	age       int
	birthYear int
}

func (p Person) Greet() {
	fmt.Printf("My name is %s I am %d years old", p.name, p.age)
}

func main() {
	fmt.Println("=========== Arrays ===========")
	// Fixed size arrays
	var arr [5]int = [5]int{12, 20, 30}
	fmt.Println(arr)

	fmt.Println("=========== Slices ===========")

	var slice []int = []int{1, 2, 3, 4, 5, 5}

	slice = append(slice, 200)
	fmt.Println("slice", slice)

	fmt.Println("=========== Maps ===========")

	dict := map[string]int{"age": 26, "birthYear": 1998}
	fmt.Println(dict["birthYear"])

	dict["age"] = 27
	fmt.Println(dict)

	fmt.Println("=========== Struct ===========")
	p := Person{"Santheep", 26, 1998}
	fmt.Println(p)
	p.Greet()

}
