package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("GoRoutine at the beginning:", runtime.NumGoroutine())
	wg.Add(2)
	go first()
	go second()
	wg.Wait()

	fmt.Println("GoRoutine at the end:", runtime.NumGoroutine())
	fmt.Println("Done!")

}

func first() {
	defer wg.Done()
	fmt.Println("GoRoutine when running first:", runtime.NumGoroutine())
	for a := 1; a <= 10; a++ {
		fmt.Println("First: ", a)
		time.Sleep(1000 * time.Millisecond)
	}
}

func second() {
	defer wg.Done()
	fmt.Println("GoRoutine when running second:", runtime.NumGoroutine())
	for a := 1; a <= 10; a++ {
		fmt.Println("Second: ", a)
		time.Sleep(100 * time.Millisecond)
	}
}
