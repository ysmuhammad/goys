package main

import (
	"fmt"
)

func main() {
	fmt.Printf("- Executing hi!\n")
	hi()

	fmt.Printf("\n- Executing odd checking!\n")
	odd()

	fmt.Printf("\n- Executing switch!\n")
	test()

	fmt.Printf("\n- Executing switch using true | false clause!\n")
	test2()

	fmt.Printf("\n- Executing switch input\n")
	ys := unhex(5)
	fmt.Printf("YS value by unhex func: %v", ys)

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
	case "beach", "hawaii", "bali":
		fmt.Println("best place to chill")
	case "mountain", "everest", "alpine":
		fmt.Println("not that good")
		fallthrough
	default:
		fmt.Println("exec this by fallthrough")
	}
}

func test2() {

	switch 2 == 2 {
	case false:
		fmt.Println("False!")
	case true:
		fmt.Println("True!")
	default:
		fmt.Println("no matching clause!")
	}
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return c
}
