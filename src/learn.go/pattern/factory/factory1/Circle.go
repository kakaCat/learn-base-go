package factory1

import "fmt"

type Circle struct {
}

func (this Circle) Draw() {
	fmt.Println("Inside Circle::draw() method.")
}
