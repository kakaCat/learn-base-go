package main

import (
	test "00learngo/learn-base-go/src/learn.go/data/base/grammar/scope/range"
	"fmt"
)

var varAll int = 1

func main() {

	var varFuc int = 1
	fmt.Println(varFuc)
	fmt.Println(varAll)
	foo()
	AllRange := test.AllRange
	AllRange()
	fmt.Println("verison")

}

func foo() {
	fmt.Println(varAll)
}
