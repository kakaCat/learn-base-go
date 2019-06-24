package factory1

import "fmt"

type Rectangle struct {
}

func (this Rectangle) Draw() {
	fmt.Println("Inside Rectangle::draw() method.")
}
