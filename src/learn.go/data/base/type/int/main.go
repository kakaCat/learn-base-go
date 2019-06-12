package main

import "fmt"

func main() {
	// int 可以是int8 int16 int32 int64
	var foo int = 9223372036854775807
	fmt.Println(foo)
	// int8   8 位整型范围-128 到 127 2^8-1
	var foo8 int8 = 127
	fmt.Println(foo8)
	// int16   16位整型范围-32768 到 32767 2^16-1
	var foo16 int16 = 32767
	fmt.Println(foo16)
	// int32   32位整型范围-2147483648 到 2147483647 2^32-1
	var foo32 int32 = 2147483647
	fmt.Println(foo32)

	// int64   64位整型范围-9223372036854775808至9223372036854775807 2^32-1
	var foo64 int64 = 9223372036854775807
	fmt.Println(foo64)

	//16进制
	var foo0x int = 0x22
	fmt.Println(foo0x)

}
