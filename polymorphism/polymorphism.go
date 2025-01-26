package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}
type Rat struct{}

func (Dog) Speak() string {
	return "Woof!"
}

func (Cat) Speak() string {
	return "Mue!"
}
func (Rat) Speak() string {
	return "Keach!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Rat{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

}
