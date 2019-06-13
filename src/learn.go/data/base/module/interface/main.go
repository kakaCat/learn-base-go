package main

import "fmt"

type Printer interface {
	Print()
}

type user struct {
	name string
	age  int
}

func (u user) Print() string {
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
