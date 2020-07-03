package main

import (
	"fmt"
)

func main() {
	a := 99
	fmt.Println(a)
	fmt.Println(&a)

	b := &a
	fmt.Println(b)

	*b = 999
	fmt.Println(a)

	// var c *string
	c := "yes"
	// fmt.Println("Value of *c:",*c)
	fmt.Println("Value of &c:", &c)
	fmt.Println("Value of c:", c)
	foo(&c)
	fmt.Println(c)
	fmt.Println(&c)

	ya := person{
		"John Wick",
		"Cilandak",
	}

	changeMe(&ya)
	fmt.Println("Name:", ya.name)
	fmt.Println("Address:", ya.address)

	kali := 2
	var kalii *int
	kalii = &kali
	*kalii = 99

	fmt.Println("kali:", kali)
	fmt.Println("kalii:", kalii)
}

func foo(s *string) {
	z := " I have been added by pointer!"
	fmt.Println("Value of *s:", *s)
	fmt.Println("Value of s:", s)
	*s = *s + z
}

type person struct {
	name    string
	address string
}

func changeMe(p *person) {
	addressChanger := "Beverly Hills"
	(*p).address = addressChanger
}
