package main

import (
	"fmt"
	"strconv"
)

type ohana int

var custom ohana
var fam string = "it means family"

func main() {
	a := "test string"
	b := 99
	c := true

	fmt.Printf("Type of a: %T\nValue of a: %v\n\n", a, a)
	fmt.Printf("Type of b: %T\nValue of b: %v\n\n", b, b)
	fmt.Printf("Type of c: %T\nValue of c: %v\n\n", c, c)

	all := fmt.Sprintf(`All Values (a,b,c): "%v","%v","%v"
All Types (a,b,c): "%T","%T","%T"`, a, b, c, a, b, c)
	fmt.Println(all)
	fmt.Println("")

	custom = 3012010
	customint := int(custom)
	us := strconv.Itoa(customint) + ", " + fam

	fmt.Println(us)

}
