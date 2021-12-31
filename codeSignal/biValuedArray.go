// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 2, 1, 3, 3, 3, 4, 4, 4, 4, 2, 1, 5}
	fmt.Println("Max adjacent range: ", solution(a))
}

func solution(arr []int) int {
	lastSeen := -1
	secondLastSeen := -1
	lbs := 0
	tempCount := 0
	lastSeenNumberRepeatedCount := 0

	for i := 0; i < len(arr); i++ {
		current := arr[i]
		if current == lastSeen || current == secondLastSeen {
			tempCount++
		} else {
			tempCount = lastSeenNumberRepeatedCount + 1
		}

		if current == lastSeen {
			lastSeenNumberRepeatedCount++
		} else {
			lastSeenNumberRepeatedCount = 1
			secondLastSeen = lastSeen
			lastSeen = current
		}
		fmt.Println(tempCount)
		lbs = Max(tempCount, lbs)
	}
	return lbs
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
