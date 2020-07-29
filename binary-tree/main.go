package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {
	fmt.Println("=== Binary Tree with channel ===")
	// First channel
	t := tree.New(1)
	ch := make(chan int)

	go Walker(t, ch)

	func() {
		fmt.Println("Channel value from tree1:")
		for i := range ch {
			fmt.Printf("%v ", i)
		}
		fmt.Println("")
	}()

	// Second chnannel
	t2 := tree.New(2)
	c2 := make(chan int)

	go Walker(t2, c2)

	func() {
		fmt.Println("Channel value from tree2:")
		for i := range c2 {
			fmt.Printf("%v ", i)
		}
		fmt.Println("")
	}()

	fmt.Println("\nSame Tree?", Same(t, t2))

}

func Walker(t *tree.Tree, ch chan int) {
	defer close(ch)
	Walk(t, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	} else if t.Left == nil {
		ch <- t.Value
		if t.Right != nil {
			Walk(t.Right, ch)
		}
		return
	} else {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	var rs bool
	c1 := make(chan int)
	c2 := make(chan int)

	go Walker(t1, c1)
	go Walker(t2, c2)

	var check1 [10]int
	var check2 [10]int

	func() {
		for z := 0; z < 10; z++ {
			check1[z] = <-c1
			check2[z] = <-c2
		}
		for y := range check1 {
			if check1[y] != check2[y] {
				rs = false
			} else {
				rs = true
			}
		}
	}()
	return rs
}
