package main

import (
	"fmt"
	"time"
)

func main() {
	go task(1)
	go task(2)

	time.Sleep(time.Second * 6)
}

//
func consumer(data chan int, done chan bool) {
	for x := range data {
		println("recv:", x)
	}

	done <- true

}

func producer(data chan int) {

	for i := 0; i < 4; i++ {
		data <- i
	}
	close(data)
}
