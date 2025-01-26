package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/santheepkumar/mathmodpublic"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("New Router Created", r)

	mathmodpublic.Greet("Santheep \n")
	fmt.Println(mathmodpublic.Add(10, 20))
	fmt.Println(mathmodpublic.Sub(30, 20))
	fmt.Println(mathmodpublic.Divide(30, 20))
}
