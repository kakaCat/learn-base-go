package main

import "fmt"

//添加类型定义
//初始化为0
var a int

//初始化为0
var b float32

//自动判读类型
var c = false

//相同类型 多个变量定义
var d, e int

//组定义类型
var (
	x, y int
	z, s = 100, "abc"
)

func main() {

	//在方法内的类型 必须被用到
	var var_int_dominant int
	fmt.Println(var_int_dominant)

	//简短模式
	//简短模式必须 在方法里 有效
	//
	var_int_recessive := 100
	fmt.Println(var_int_recessive)

	//先运算，再赋值
	x_1, y_1 := 1, 2
	x_1, y_1 = y_1+3, x_1+2
	fmt.Println(x_1, y_1)
	x_1, y_1 = y_1+x_1, x_1+y_1
	fmt.Println(x_1, y_1)

}
