package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Weight float32
}

func (p Person) Greet() {
	fmt.Printf("Welcome %s", p.Name)
}

func (p *Person) ChangeName(name string) {
	p.Name = name
	fmt.Printf("Name updated %s", p.Name)
}

func main() {
	p := Person{Name: "Santheep", Age: 26, Weight: 70}
	p.Age = 30
	// fmt.Println("Person Struct", p.Age)
	p.Greet()
	p.ChangeName("Rahul \n")
	p.Greet()
}
