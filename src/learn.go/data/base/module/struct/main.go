package main

import "fmt"

type user struct {
	name string
	age  int
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
	fmt.Println("manager :", m.ToString())

}
