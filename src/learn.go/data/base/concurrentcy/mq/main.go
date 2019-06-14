package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

func main() {
	//channl_default_1()
	//channl_default_2()
	//channl_default_3()
	//channl_default_4()
	//channl_default_5()
	//channl_default_6()
	channl_default_7()
}

func channl_default() {
	//结束事件
	done := make(chan struct{})
	//数据传输通道
	c := make(chan string)
	go func() {
		//接收消息
		s := <-c
		println(s)
		//关闭管道，作为通知
		close(done)
	}()

	c <- "hi!"
	//阻塞，直到有数据或管道关闭
	<-done
}

func channl_default_1() {
	//创建带3个缓冲槽的异步通道
	c := make(chan int, 3)
	//缓存区未满，不会阻塞
	c <- 1
	c <- 2
	println(<-c)
	println(<-c)
}

func channl_default_2() {
	//创建带3个缓冲槽的异步通道
	c := make(chan int, 3)
	//缓存区未满，不会阻塞
	c <- 1
	c <- 2
	println(<-c)
	println(<-c)
}

func channl_default_3() {
	var a, b chan int = make(chan int, 3), make(chan int)
	var c chan bool
	println(a == b)
	println(c == nil)

	fmt.Println("%p, %d\n", a, unsafe.Sizeof(a))
}

func channl_default_4() {
	var a, b chan int = make(chan int), make(chan int, 3)
	b <- 1
	b <- 2
	println("a:", len(a), cap(a))
	println("b:", len(b), cap(b))
}

func channl_default_5() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)
		for {
			x, ok := <-c
			if !ok {
				return
			}
			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	close(c)
	<-done
}

func channl_default_6() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for x := range c {
			println(x)
		}

	}()

	c <- 1
	c <- 2
	c <- 3
	close(c)
	<-done
}

func channl_default_7() {

	var wg sync.WaitGroup
	ready := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(i)
		go func(id int) {
			defer wg.Done()

			println(id, ": ready.")
			<-ready
			println(id, ": runing.")
		}(i)
		time.Sleep(time.Second)
		println("Ready? Go!")
		close(ready)
		wg.Wait()
	}

}
