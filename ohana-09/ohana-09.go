package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByName []Person

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Less(i, j int) bool {
	return a[i].First < a[j].First
}

func main() {
	p1 := Person{
		"Z",
		33,
	}
	p2 := Person{
		"Q",
		28,
	}

	people := []Person{p1, p2}
	fmt.Println(people)

	sort.Sort(ByAge(people))
	fmt.Println(people)

	people2 := []Person{
		Person{
			"D",
			40,
		},
		Person{
			"E",
			50,
		},
		Person{
			"B",
			20,
		},
		Person{
			"C",
			30,
		},
		Person{
			"A",
			10,
		},
	}

	cek := ByName(people2).Less(0, 4)
	fmt.Printf("Check fact of %s is less than %s: %v\n", people2[0].First, people2[4].First, cek)

}
