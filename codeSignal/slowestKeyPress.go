package main

import "fmt"

func main() {
	a := [][]int{{0, 2}, {1, 5}, {0, 9}, {2, 15}}
	fmt.Println(sol(a))
}

func sol(arr [][]int) string {
	m := make(map[int]int)
	var max, highestVal int
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			m[arr[i][0]] = arr[i][1]
			continue
		}
		if _, ok := m[arr[i][0]]; !ok {
			m[arr[i][0]] = arr[i][1] - arr[i-1][1]
		} else {
			if m[arr[i][0]] < arr[i][1]-arr[i-1][1] {
				m[arr[i][0]] = arr[i][1] - arr[i-1][1]
			}
		}
	}
	fmt.Println(m)
	for key, val := range m {
		if val > max {
			max = val
			highestVal = key
		}
	}
	return string(highestVal + 97)
}

