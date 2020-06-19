package main

import (
	"fmt"
)

func foo () int{
	return 53
}

func bar () (string, int){
	bs := "Hello!"
	bi := 99
	return bs, bi
}

type person struct{
	first string
	last string
	age string
}

func speak (p person) {
	fmt.Println("Hi my name is:", p.first, p.last)
	fmt.Println("Age:",p.age)
}

func main (){
	s, i := bar()
	fmt.Println(s,i)
	ys := person{
		first: "Yusuf",
		last: "Muhammad",
		age: "23",
	}
	speak(ys)
}