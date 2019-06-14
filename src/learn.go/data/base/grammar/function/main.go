package main

import "fmt"

func main() {

	a, b := 1, 2
	c := add(a, b)
	fmt.Println("add result:", c)

	d, e := add_error(a, b)
	fmt.Println("add_error result:", d, e)

	f, g, h := returnMoreDate3(a)
	fmt.Println("returnMoreDate3 result:", f, g, h)

	i := func_defer(a)
	fmt.Println("func_defer result:", i)

	l := func_defer_more(a, b)
	fmt.Println("func_defer_more result:", l)
}

//方法没有重载
func add(a, b int) int {

	return a + b
}

//func add(a, b,c int) int {
//
//	return a + b
//}

func add_error(a, b int) (int, error) {

	return a + b, nil
}

func returnMoreDate3(a int) (int, int, string) {
	return a, 1, "ss"
}

//关键字defer 函数结束后执行 用于释放资源，解锁 执行一些清理操作
func func_defer(a int) int {
	defer fmt.Println("defer end result:", a)

	fmt.Println("defer start result:", a)

	return a
}

//关键字defer 多个出现的时候 从上到下执行
func func_defer_more(a, b int) int {
	defer fmt.Println("defer a end result:", a)
	defer fmt.Println("defer b end result:", b)

	fmt.Println("defer start result:", a)

	return a
}

func func_anonymity(s string) {

}
