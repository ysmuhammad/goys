package main

import (
	"fmt"
)

func foo() int {
	return 53
}

func bar() (string, int) {
	bs := "Hello!"
	bi := 99
	return bs, bi
}

type person struct {
	first string
	last  string
	age   string
}

func speak(p person) {
	fmt.Println("Hi my name is:", p.first, p.last)
	fmt.Println("Age:", p.age)
}

type circle struct {
	radius float64
}

type square struct {
	panjang float64
	lebar   float64
}

func (s *square) area() float64 {
	luas := s.panjang * s.lebar
	return luas
}

func (c *circle) area() float64 {
	const pi = 3.14

	luas := pi * c.radius * c.radius
	return luas
}

type shape interface {
	area() float64
}

func info(s shape) string {
	var sVal string
	sVal = fmt.Sprintf("Total Area: %#v", s.area())
	return sVal
}

func main() {
	s, i := bar()
	fmt.Println(s, i)
	ys := person{
		first: "Yusuf",
		last:  "Muhammad",
		age:   "23",
	}
	speak(ys)

	fmt.Println("========================")

	cirval := circle{
		radius: 16,
	}
	v1 := cirval.area()
	fmt.Println("Circle area: ", v1)

	sqval := square{
		panjang: 30,
		lebar:   40,
	}
	v2 := sqval.area()
	fmt.Println("Square area: ", v2)

	v3 := info(&cirval)
	fmt.Println(v3)

	fmt.Println("========================")

	var test *string
	test2 := "Test value!"
	test = &test2
	fmt.Printf("Value of test: %v", *test)

}
