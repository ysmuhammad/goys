package main

import (
	"fmt"
)

func main() {
	fmt.Printf("executing hi!\n")
	hi()

	fmt.Printf("\nexecuting odd checking!\n")
	odd()

	fmt.Printf("\nexecuting switch!\n")
	test()
}

func hi() {
	hi := 1
	for {
		if hi > 9 {
			break
		}
		if hi > 1 {
			fmt.Printf("Hi! %v times\n", hi)
		} else {
			fmt.Printf("Hi! %v time\n", hi)
		}
		hi += 1
	}
}

func odd() {
	odd := 1
	for {
		odd += 1

		if odd%6 == 0 {
			continue
		}

		if odd%3 == 0 {
			fmt.Println(odd)
		}

		if odd > 33 {
			break
		}
	}
}

func test() {

	s := "beach"

	switch s {
	case "beach":
		fmt.Println("best place to chill")
	case "mountain":
		fmt.Println("not that good")
		fallthrough
	default:
		fmt.Println("exec this by fallthrough")
	}
}
