package main

import "fmt"

func main() {

	x := 1
	ifDome(x)
	switchDome_fallthrough(x)

}

func ifDome(x int) {
	if x > 0 {
		fmt.Println(">0", x)
	} else if x < 0 {
		fmt.Println("<0", x)
	} else {
		fmt.Println("else", x)
	}
}

func switchDome(x int) {

	switch {
	case x > 0:
		fmt.Println(">0", x)
	case x < 0:
		fmt.Println("<0", x)
	default:
		fmt.Println("default", x)
	}

}

//fallthrough关键字 继续执行 break
func switchDome_fallthrough(x int) {

	switch {
	case x > 0:
		fmt.Println("fallthrough >=1", x)
		fallthrough
	case x < 0:
		fmt.Println("<=1", x)
		if 1 == x {
			break
		}
		fallthrough
	default:
		fmt.Println("default", x)
	}
}
