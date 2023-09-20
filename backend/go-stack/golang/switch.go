package main

import "fmt"

func main() {
	i := 6
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4, 5:
		fmt.Println("three, four, five")
	default:
		fmt.Println("bruv")
	}
}
