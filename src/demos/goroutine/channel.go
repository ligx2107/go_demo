package gr

import "fmt"

/*
	channel:
		1. channel是goroutine之间的连接，是goroutine之间的一种通信机制，可以让一个goroutine发送特定的值到另外一个goroutine
		2. channel是Go语言的一种特殊类型，像队列一样遵守FIFO原则，从而保证收发数据的顺序
		3. 每个channel都是一个具体类型的管道，在声明channel时需要为其指定元素类型
		4. 声明channel类型的变量： var 变量 chan 元素类型
		5. 创建channel的格式：make(chan 元素类型, [缓冲大小])
		6. channel的基本操作：
			6.1 发送： ch <- 数据
			6.2 接收： x := <- ch
			6.3 关闭： close(ch)
*/

func TestChannel() {
	// 创建一个channel
	cha := make(chan int)

	// 创建goroutine
	go func() {
		defer fmt.Println("coroutine end...")
		fmt.Println("coroutine is running...")
		// 向管道cha中写入数据
		cha <- 666
	}()

	// main coroutine从channel中获取数据
	data := <-cha
	fmt.Println("main coroutine get data -->", data)
}
