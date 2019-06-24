package main

import (
	"fmt"
	"reflect"
)

type X int
type Y int

func main() {
	method2()
}

func method1() {
	var a X = 100
	t := reflect.TypeOf(a)
	fmt.Println("t.Name", t.Kind())
}

//Type是真实类型(静态类型)
//Kind是基础类型(底层类型)
func method2() {
	var a, b X = 100, 200
	var c Y = 300
	ta, tb, tc := reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c)
	fmt.Println("ta:tb", ta == tb)
	fmt.Println("ta:tc", ta == tc)
	fmt.Println("ta:tc kind", ta.Kind() == tc.Kind())

}
