package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100

	b := [5]int{1, 3, 5, 6, 7}
	fmt.Println(b)
}
