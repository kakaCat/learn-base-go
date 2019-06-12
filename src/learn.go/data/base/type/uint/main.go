package main

import "fmt"

func main() {
	// uint 可以是uint8 uint16 uint32 uint64 自动转换
	var foo uint = 18446744073709551615
	fmt.Println(foo)
	// uint8   8 位整型范围0 到 255   2^8*2-1
	var foo8 uint8 = 255
	fmt.Println(foo8)
	// uint16   16位整型范围0 到 65535 2^16*2-1
	var foo16 uint16 = 65535
	fmt.Println(foo16)
	// uint32   32位整型范围0 到 4294967295 2^32*2-1
	var foo32 uint32 = 4294967295
	fmt.Println(foo32)

	// uint64   64位整型范围0 到 18446744073709551615 2^32*2-1
	var foo64 uint64 = 18446744073709551615
	fmt.Println(foo64)
}
