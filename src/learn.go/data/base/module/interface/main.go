package main

import "fmt"

//不能有字段
//不能定义自己的方法
//只能声明方法，不能实现
//可嵌入其他接口类型
type Printer interface {
	Print()
}

type user struct {
	name string
	age  int
}

type user2 struct {
	name string
	age  int
}

func (u user) Print() string {
	return fmt.Sprintf("%+v", u)
}

func (u user2) Print() string {
	return fmt.Sprintf("%+v", u)
}

type manager struct {
	user
	title string
}

func (m manager) ToString() string {
	return fmt.Sprintf("%+v", m)
}

func main() {
	var m manager
	m.name = "tom"
	m.age = 29
	m.title = "test"
	fmt.Println("manager :", m.Print())

}
