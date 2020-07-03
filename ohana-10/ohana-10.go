package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)

	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("CPU:", runtime.NumCPU())
	gorut := runtime.NumGoroutine()
	fmt.Println("GoRoutine:", gorut)
	fmt.Println("Arch:", runtime.GOARCH)

	fmt.Println("--------------")

	var wg sync.WaitGroup

	wg.Add(2)

	fmt.Println("Starting Go Routines")
	go func() {
		defer wg.Done()

		for char := 'a'; char < 'a'+999; char++ {
			fmt.Printf("%c ", char)
		}
	}()

	go func() {
		defer wg.Done()

		for number := 1; number < 999; number++ {
			fmt.Printf("%d ", number)
		}
	}()

	fmt.Println("Waiting To Finish")
	gorut = runtime.NumGoroutine()
	fmt.Println("GoRoutine:", gorut)
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
