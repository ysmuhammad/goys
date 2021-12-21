package main

import "fmt"

func solution(inputString string) bool {
	l := len(inputString) - 1
	lastWord := string(inputString[l])
	vocal := []string{"a", "i", "u", "e", "o"}
	for _, v := range vocal {
		if v == lastWord {
			return true
		}
	}
	return false
}

func main() {
	ans := solution("yeww")
	fmt.Println(ans)
}
