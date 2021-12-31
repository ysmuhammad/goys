package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	arr := []int{25, 35, 872, 228, 53, 278, 872}
	fmt.Println("Digit Anagram: ", solution(arr))
}

func solution(arr []int) int {
	var s [][]string
	var count int
	for i := 0; i < len(arr); i++ {
		var val []string
		val = strings.Split(strconv.Itoa(arr[i]), "")
		sort.Strings(val)
		s = append(s, val)
	}
	for k := 0; k < len(s)-1; k++ {
		for l := k + 1; l < len(s); l++ {
			if testEq(s[k], s[l]) {
				count += 1
			}
		}
	}

	return count
}

func testEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
