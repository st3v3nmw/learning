package main

import "fmt"

func main() {
	nums := []int{3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "Stephen" {
		fmt.Println(i, c)
	}
}
