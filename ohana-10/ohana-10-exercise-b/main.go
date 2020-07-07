package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	//runtime.GOMAXPROCS(2)
	var counter int64
	counter = 0

	const gs = 10
	var wg sync.WaitGroup
	wg.Add(gs)

	var i int64
	for i = 0; i < gs; i++ {
		go func(i int64) {
			atomic.AddInt64(&counter, i)
			// v := counter
			// v++
			// counter = v
			atomic.LoadInt64(&counter)
			wg.Done()
		}(i)
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
