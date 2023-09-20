package main

import "fmt"

type rect struct {
	width, height uint
}

func (r *rect) area() uint {
	return r.width * r.height
}

func (r rect) perim() uint {
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area", r.area())
	fmt.Println("perimeter", r.perim())
}
