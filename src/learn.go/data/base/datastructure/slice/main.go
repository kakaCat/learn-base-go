package main

import "fmt"

func main() {
	slice_dome()
}

func slice_dome() {
	//初始化 长度为5
	x := make([]int, 0, 5)
	for i := 0; i < 8; i++ {
		x = append(x, i)
	}
	fmt.Println("append", x)
}
