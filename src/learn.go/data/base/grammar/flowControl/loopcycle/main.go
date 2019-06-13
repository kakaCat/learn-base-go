package main

import "fmt"

func main() {
	fordome()
	for_range()
}

func fordome() {
	for i := 0; i < 5; i++ {
		fmt.Println("for", i)
	}
}

func for_range() {

	x := [3]int{1, 2, 3}
	for i, n := range x {
		fmt.Println("range", i, ":", n)
	}
}
