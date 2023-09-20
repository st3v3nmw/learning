package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 3
	m["k2"] = 4
	fmt.Println("map:", m)

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	m1 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", m1, maps.Equal(m, m1))
}
