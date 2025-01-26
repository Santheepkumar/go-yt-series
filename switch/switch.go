package main

import "fmt"

func main() {
	day := "Sunday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of the week")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Regular day")
	}

	score := 12

	switch {
	case score >= 100:
		fmt.Println("Grade S")
	case score >= 50:
		fmt.Println("Grade A")
	default:
		fmt.Println("Grade E")
	}

}
