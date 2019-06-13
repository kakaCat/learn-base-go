package main

import "fmt"

func main() {
	map_dome()
}

//
func map_dome() {
	//初始化 长度为5
	x := make(map[string]int)
	//添加
	x["a"] = 1
	//获取
	a, ok := x["a"]
	fmt.Println("获取a:", a, ok)
	b, ok := x["b"]
	fmt.Println("获取b:", b, ok)
	//删除
	delete(x, "a")
	fmt.Println("获取a 是指针还是值:", a, ok)

	aa, ok := x["a"]
	fmt.Println("获取a:", aa, ok)
}
