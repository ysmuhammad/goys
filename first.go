package main

import (
	"fmt"
	"strconv"
)

var ys int

var dv string

func main() {
	ys = 6
	dv = "alana"

	fmt.Println(ys, dv)
	dv = strconv.Itoa(ys)
	fmt.Println(dv)
	fmt.Printf("%T\n", dv)
}
