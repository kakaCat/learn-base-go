package main

import "fmt"

//常量只读，不能修改
const x, y int = 123, 0x22
const s = "hello world"
const c = '我'

const (
	i, f = 1, 0.123
	b    = false
)

func main() {

	fmt.Println("global x", x)
	const x = 124
	fmt.Println("local x", x)

	fmt.Println("global i", i)
}
