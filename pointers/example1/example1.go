package main
import "fmt"

func main() {
    x := 10
    p := &x  

    fmt.Println("Address stored in p:", p)
    fmt.Println("Value at address p:", *p) 

    *p = 20 

    fmt.Println("New value of x:", x) 
}
