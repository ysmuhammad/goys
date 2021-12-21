package main

import (
	"fmt"
)

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

	printFunc := returnFunc()("Assalamualaikum Wr Wb!")
	fmt.Printf("\n%v\n", printFunc)

	fmt.Println("\nCallback function integer:")
	val := []int{1, 2, 3, 4, 5}
	callbackFunc := even(sum, val...)
	fmt.Println(callbackFunc)

	fmt.Println("Callback function string:")
	wo := []string{"what", "the", "fuck", "is", "going", "on", "?"}
	p := wordsUp(wo...)
	fmt.Println("No filter:", p)

	callbackFunc2 := goodWords(wordsUp, wo...)
	fmt.Println("Only good words:", callbackFunc2)

	recur := rec(4)
	fmt.Println("Print recursive:", recur)

	test := "abcde"
	fmt.Println(string(test[4]))
}

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

// Returning a function
func returnFunc() func(y string) string {
	return func(x string) string {
		yesFunc := x
		return yesFunc
	}
}

// Callback function
// Sum func
func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(i ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	totalEven := f(yi...)
	return totalEven
}

func wordsUp(w ...string) string {
	var allWords string
	for i, v := range w {
		if i == 0 {
			allWords = allWords + v
		} else {
			allWords = allWords + " " + v
		}
	}
	return allWords
}

func goodWords(f func(w ...string) string, si ...string) string {
	var goodWordsArr []string
	for _, val := range si {
		if val == "fuck" {
			continue
		} else if val == "the" {
			continue
		} else {
			goodWordsArr = append(goodWordsArr, val)
		}
	}
	return f(goodWordsArr...)
}

//Recursive
func rec(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
