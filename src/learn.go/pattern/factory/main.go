package main

import "00learngo/learn-base-go/src/learn.go/pattern/factory/factory1"

func main() {
	factory := factory1.ShapeFactory{}
	factory.GetShape("1").Draw()
	factory.GetShape("2").Draw()
	factory.GetShape("3").Draw()
}
