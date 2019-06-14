package main

import (
	"sync"
	"time"
)

func main() {
	wait_default_sync()
}

func wait_default() {
	exit := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		println("goroutine done .")
		close(exit)
	}()
	println("main ...")

	<-exit
	println("main exit .")
}

func wait_default_sync() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		//累积计数
		wg.Add(i)
		go func(id int) {
			//递减计数
			defer wg.Done()

			time.Sleep(time.Second)
			println("goroutine:", id, "done.")
		}(i)
	}

	println("main ...")

	wg.Wait()
	println("main exit .")
}
