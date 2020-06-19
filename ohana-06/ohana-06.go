package main

import (
	"fmt"
)

type players struct {
	firstName string
	lastName  string
	age       int
	cars
}

type cars struct {
	tipe  string
	color string
}

func main() {
	xi := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	fmt.Println(xi)

	dv, ys := sl("alana", xi...)
	fmt.Printf("Momma: %v\nTotal: %v", dv, ys)

	// Anon function
	gold := func(x int) {
		fmt.Println("Gold number:", x)
	}
	gold(99)

	roo := players{
		firstName: "Rooney",
		lastName:  "Wayne",
		age:       30,
		cars: cars{
			tipe:  "Mercedes Benz",
			color: "Red",
		},
	}

	merci := cars{
		tipe:  "Mercedez Benz",
		color: "White",
	}

	merci.score()
	roo.score()

	champ(roo)
	champ(merci)

	returnFunc()
	fmt.Printf("\n%v\n", returnFunc()())

	// v := rt()
	// fmt.Println(v)

}

func sl(y string, x ...int) (int, string) {

	defer fmt.Println("End of sl function!\n==============")
	ys := 0
	for _, v := range x {
		fmt.Println(v)
		ys += v
	}
	tmpAlana := y
	anuy := fmt.Sprintf("Binary: %b \n\n", byte(tmpAlana[0]))
	fmt.Printf("\nType of anuy: %T\nValue of anuy: %v", anuy, anuy)

	return ys, anuy
}

func (s players) score() {
	fmt.Printf("Scorer: %v\n", s.firstName)
	fmt.Printf("Age: %v\n", s.age)
	fmt.Printf("Car: %v\n", s.tipe)
	fmt.Printf("Car color: %v\n", s.color)
}

func (s cars) score() {
	fmt.Printf("Car: %v\n", s.tipe)
	fmt.Printf("Car color: %v\n", s.color)
}

type champions interface {
	score()
}

func champ(c champions) {
	switch c.(type) {
	case players:
		fmt.Printf("\nYes you are the champion, %v %v!", c.(players).lastName, c.(players).firstName)
	case cars:
		fmt.Printf("\nBest car out there is: %v %v\n", c.(cars).color, c.(cars).tipe)
	}
}

func returnFunc() func() string {
	return func() string {
		return "End of!"
	}
}
