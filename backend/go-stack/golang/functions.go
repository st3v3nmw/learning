package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func vals() (int, int) {
	return 4, 6
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("1+2 =", plus(1, 2))
	b, _ := vals()
	fmt.Println(b)
	fmt.Println(sum(1, 3, 4), sum(1))
}
