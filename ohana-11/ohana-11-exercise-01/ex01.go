package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go receive(c, done)
	gen(c)

	fmt.Println("about to exit")
}

func gen(c chan int) {
	defer close(c)
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func receive(ch chan int, done chan bool) {
	for {
		val, i := <-ch
		if i {
			fmt.Println("Value:", val)
		} else {
			fmt.Println("All done!")
			done <- true
			return
		}
	}
}