package gr

import (
	"fmt"
	"time"
)

/*
	select多路复用：同时监听一个或多个channel，直到其中一个channel ready，如果多个channel同时ready，则随机选择一个执行，所有channel都不可用时，执行default语句
	select {
		case <- chan1:
			如果chan1成功读取数据，则执行该case处理语句
		case chan1 <- 1:
			如果chan1成功写入数据，则执行该case处理语句
		default:
			如果以上都没有成功，则执行default处理语句
	}
*/

func test(cha1 chan<- int, cha2 <-chan int) {
	x := 1
	for {
		// select监听两个channel
		select {
		case cha1 <- x:
			// 如果可以向cha1中写入数据，则进行如下计算
			x++
		case <-cha2:
			// 如果可以从cha2中读取数据，则进行如下处理
			fmt.Println("end...")
			return
		}
	}
}

func TestSelect() {
	// 创建两个channel
	cha1 := make(chan int)
	cha2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-cha1)
		}
		cha2 <- 0
	}()

	go test(cha1, cha2)

	time.Sleep(2 * time.Second)
}
