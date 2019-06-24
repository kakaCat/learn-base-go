package factory1

import "fmt"

type Square struct {
}

func (this Square) Draw() {
	fmt.Println("Inside Square::draw() method.")
}
