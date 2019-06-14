package main

import "fmt"

type N int

func (n N) toString() string {
	return fmt.Sprint("%#x", n)
}

func (n N) value() {
	n++
	fmt.Printf("v: %p ,%v \n", &n, n)
}

func (n *N) prointer() {
	(*n)++
	fmt.Printf("p: %p ,%v \n", n, *n)
}
func prointer_default() {
	var a N = 25
	println(a.toString())
}

func prointer_default_1() {
	var a N = 25
	a.value()
	a.prointer()
	fmt.Printf("a: %p, %v\n", &a, a)
}

func prointer_default_2() {
	var a N = 25
	p := &a
	a.value()
	a.prointer()
	p.value()
	p.prointer()

}

func main() {
	prointer_default_1()
	fmt.Println("-----------")
	prointer_default_2()

}
