package main

import "fmt"

func main() {
	var s []string
	fmt.Println("uninint:", s, s == nil, len(s))

	s = make([]string, 3)
	s = append(s, "d", "e")
	fmt.Println("emp:", s, cap(s), len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	fmt.Println(c[2:5])
}
