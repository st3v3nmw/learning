package main

import "fmt"

func max[T int | float64](nums ...T) (result T) {
	result = nums[0]
	for _, num := range nums[1:] {
		if num > result {
			result = num
		}
	}
	return result
}

func main() {
	fmt.Println(max(0, -4, 7, 1, 9))
}
