package main

import "fmt"

func main() {
	var s = "Go🔥"
	banner(s, 6)

	for i, r := range s {
		fmt.Println(i, r)
		fmt.Printf("%c of type %T\n", r, r)
	}

	x, y := 1, "2"
	fmt.Printf("x=%#v, y=%#v\n", x, y)

	fmt.Println(isPalindrome("dog"))
	fmt.Println(isPalindrome("dod"))
	fmt.Println(isPalindrome("a"))
}

func banner(text string, width int) {
	padding := (width - len(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}
