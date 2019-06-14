package main

import "fmt"

func main() {
	arr_default()

}

//未初始化的默认初始化
func arr_default() {
	var a [4]int
	b := [4]int{2, 5}
	//指定位置初始化
	c := [4]int{5, 3: 10}
	//按照初始化值的数量定义数组长度
	d := [...]int{1, 2, 3}

	e := [...]int{10, 3: 100}

	fmt.Println(a, b, c, d, e)

}

//implicit explicit
func arr_no_length() {
	// 长度不同 不属于同一类型
	//var arr1 [3]int
	//var arr2 [2]int
	//arr1 = arr2
}
