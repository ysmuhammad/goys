package main

import (
	"fmt"
	"strings"
)

// Accum("test")
func Accum(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		result = result + strings.ToUpper(string(s[i]))
		for j := 0; j < i; j++ {
			result = result + string(s[i])
		}
		fmt.Println(result)
		if i != len(s)-1 {
			result = result + "-"
		}
	}
	return result
}

func main() {
	st := Accum("abcd")
	fmt.Println(st)
}
