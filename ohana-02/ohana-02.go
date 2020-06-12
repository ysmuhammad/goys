package main

import (
	"fmt"
)

const (
	_    = iota
	x    = 1 << (iota * 10)
	y    = 1 << (iota * 10)
	year = 2017 + iota
)

func main() {
	s := "Hello there! How you doin?"
	fmt.Println(s)

	slice := []rune(s)
	fmt.Println(slice)

	// ys := range s
	// fmt.Println(ys)

	ys := []string{"hell", "o", "world"}
	fmt.Println(ys)

	for i, n := range ys {
		fmt.Printf("Di index-%d ada ini %v\n", i, n)
	}

	fmt.Printf("\n\nDecimal of x: %d, binary of x: %b\n", x, x)
	fmt.Printf("Decimal of y: %d, binary of y: %b\n\n", y, y)

	xy := x * y
	cal := fmt.Sprintf("The calculation of x and y: %v, processed on the year of: %v", xy, year)

	fmt.Println(cal)

}
