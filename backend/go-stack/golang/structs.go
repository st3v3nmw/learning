package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base // struct embedding
	str  string
}

func main() {
	fmt.Println(person{"Bob", 20}, person{}, newPerson("revolution"))

	co := container{
		base: base{num: 1},
		str:  "some name",
	}
	fmt.Println(co.describe())

	type describer interface {
		describe() string
	}
	var d describer = co
	fmt.Println("describer:", d.describe())
}
