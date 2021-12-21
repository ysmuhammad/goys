package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x [5]int
	fmt.Println(x)
	fmt.Println(len(x))

	y := [5]string{"one", "two"}
	fmt.Println(y[1])
	fmt.Println(len(y))

	var twoD [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for z := 0; z < 4; z++ {
				twoD[i][j] = i + j + z
			}
		}
	}
	fmt.Println("2d array: ", twoD)
	fmt.Println("2d [2][0]: ", twoD[2][0])

	var threeD [3][3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for z := 0; z < 3; z++ {
				threeD[i][j][z] = i + j + z
			}
		}
	}
	fmt.Println("3d array: ", threeD)
	fmt.Println("3d [0][2][0]]: ", threeD[0][2][0])

	val := test_append(5)

	fmt.Println(val)

	test_slice()

	testMap()

	yo := 2 + 1
	fmt.Printf("\n\nType of Z: %T \n", yo)

}

func test_append(ys int) int {
	x := []int{1, 2, 3, 4, 5}
	fmt.Printf("\nX slice: %v\n", x)

	y := []int{6, 7, 8, 9}
	fmt.Printf("Y slice: %v\n", y)

	x = append(x, y...)
	x = append(x, 10, 11, 12, 13, 14, 15)
	fmt.Printf("Final X: %v\n", x)

	x = append(x[:2], x[4:]...)
	fmt.Printf("After delete some index on X: %v\n\n", x)

	ys += x[5]

	return ys

}

func test_slice() {
	dv := make([]int, 5, 15)

	num := []int{rand.Intn(10), rand.Intn(30), rand.Intn(20), rand.Intn(50), rand.Intn(80)}
	for i, ys := range num {
		dv[i] = ys
		i += 1
	}
	fmt.Printf("dv: %v\n\n", dv)
}

func testMap() {
	map1 := map[string]string{
		"Hawaii":    "Honolulu",
		"Indonesia": "Bali",
		"Thailand":  "Bangkok",
	}
	fmt.Printf("Best city on Hawaii? %v\n\n", map1["Hawaii"])

	y, ok := map1["Indonesia"]
	fmt.Printf("Type of y: %T\nType of ok: %T\n\n", y, ok)
	fmt.Printf("Value of y: %v\nValue of ok: %v\n", y, ok)

	fmt.Println("\n============")

	map1["Hawaii"] = "Oahu"

	for k, v := range map1 {
		fmt.Printf("\nKey: %v\n", k)
		fmt.Printf("Key: %v\n", v)
	}
}
