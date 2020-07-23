package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	c := make(chan int)

	go dv(c)
	go ys(c)

	wg.Wait()
}

// sending channel
func ys(c chan<- int) {
	for i := 1; i <= 100; i++ {
		c <- i
	}
	close(c)
	// fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	wg.Done()
}

//receiving channels
func dv(c <-chan int) {
	for i := range c {
		if i%2 == 0 {
			fmt.Println("Channel Value: ", i)
		}
	}
	// fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	wg.Done()
}
