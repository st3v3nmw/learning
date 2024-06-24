package main

import (
	"fmt"
	"slices"
	"sort"
)

func concat(s1, s2 []string) []string {
	return append(s1, s2...)
}

func median(vs []float64) float64 {
	values := make([]float64, len(vs))
	copy(values, vs)
	sort.Float64s(values)

	i := len(values) / 2
	if len(values)%2 == 1 {
		return values[i]
	}

	return (values[i-1] + values[i]) / 2
}

func main() {
	concated := concat([]string{"A", "B"}, []string{"C", "D", "E"})
	fmt.Println(concated)

	vs := []float64{2, 1, 3, 4}
	fmt.Println(median(vs), vs)

	s := []int{0, 1, 2, 3, 4}
	s = append(s[:2], s[3:]...)
	fmt.Println(s)

	slices.Reverse(s)
	fmt.Println(s)
}
