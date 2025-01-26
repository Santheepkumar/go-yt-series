package main

import "fmt"

type Greeter interface {
	Greet()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Welcome %s", p.Name)
}

func main() {

	var g Greeter = Person{Name: "Santheep", Age: 26}
	g.Greet()

}
