package main

import "fmt"

func main() {
	//slice_dome()
	slice_default()
}

func slice_dome() {
	//初始化 长度为5
	x := make([]int, 0, 5)
	for i := 0; i < 8; i++ {
		x = append(x, i)
	}
	fmt.Println("append", x)
}

func slice_default() {
	s1 := make([]int, 3, 5)
	s2 := make([]int, 3)
	s3 := []int{10, 20, 5: 30}
	fmt.Println(s1, len(s1), cap(s2))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
}
