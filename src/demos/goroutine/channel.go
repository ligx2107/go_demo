package gr

import (
	"fmt"
	"time"
)

/*
	channel:
		1. channel是goroutine之间的连接，是goroutine之间的一种通信机制，可以让一个goroutine发送特定的值到另外一个goroutine
		2. channel是Go语言的一种特殊类型，一个并发安全的队列，遵守FIFO原则，从而保证收发数据的顺序
		3. 每个channel都是一个具体类型的管道，在声明channel时需要为其指定元素类型
		4. 声明channel类型的变量： var 变量 chan 元素类型
		5. 创建channel的格式：make(chan 元素类型, [缓冲大小])
		6. channel的基本操作：
			6.1 发送： ch <- 数据
			6.2 接收： x := <- ch
			6.3 关闭： close(ch)，如果不在向channel内写入数据或从channel中读取数据，关闭channel
				6.3.1 读取数据goroutine判断channel是否关闭的方式： (data,ok := <- ch, ok为false时，表示channel以关闭)，(使用for range方式)
				6.3.2 关闭channel后，再向channel中写入数据会引发panic
				6.3.3 关闭channel后，再从channel中读取数据不会引发panic，读取到的数据为channel类型的零值或已缓存的数据
				6.3.4 重复关闭或关闭nil channel会引发panic
		7. 无缓冲channel：又称为阻塞channel，当没有接收者时，发送操作将阻塞；当没有发送者时，接收操作将阻塞
		8. 有缓冲channel：异步channel，当缓冲区满时，发送操作将阻塞；当缓冲区空时，接收操作将阻塞，有助于提高程序的性能，减少排队阻塞
		9. 单向channel，
			9.1 写channel：chan<- 元素类型
			9.2 读channel：<-chan 元素类型
			9.3 双向channel类型变量可以给单向channel类型变量赋值，反之则不可以
			9.4 通常使用类型转换来获取单向channel类型变量
				c := make(chan int)
				var send chan<- int = c
				var recv <-chan int = c
			9.5 不能在单向通道上做逆向操作，同时close只能关闭写通道
*/

func testNobufferChannel() {
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

func testBufferChannel(bufferCap int) {
	// 创建有缓冲channel，缓冲区大小为3
	cha := make(chan int, bufferCap)

	// 打印channel buffer的长度和容量
	fmt.Println("before len(cha) -->", len(cha), ", cap(cha) -->", cap(cha))

	// 创建goroutine
	go func() {
		defer fmt.Println("sub goroutine end...")
		// 循环向channel中发送数据
		for i := 0; i < bufferCap; i++ {
			cha <- i
			fmt.Println("len(cha) -->", len(cha), "cap(cha) -->", cap(cha))
		}
	}()

	// main goroutine睡两秒，保证子goroutine执行完毕
	time.Sleep(2 * time.Second)

	// main goroutine从channel中获取数据
	for i := 0; i < bufferCap; i++ {
		num := <-cha
		fmt.Println("num =", num)
	}

	fmt.Println("main goroutine end ...")
}

func testCloseChannel() {
	cha := make(chan int)
	go func() {
		for i := 0; i < 4; i++ {
			cha <- i
		}
		// 写数据完成，关闭channel
		close(cha)
	}()
	// 循环取数据
	for {
		// 持有初始化语句的if，可定义代码块局部变量
		if data, ok := <-cha; ok {
			fmt.Println("data -->", data)
		} else {
			// channel关闭，退出循环
			break
		}
	}

}

// cha 写channel
func writeData(cha chan<- int) {
	// 向写channel中写入数据
	defer fmt.Println("写数据结束...")
	for i := 0; i < 4; i++ {
		cha <- i
	}
	close(cha)
}

// cha 读channel
func readData(cha <-chan int) {
	for data := range cha {
		fmt.Println("data -->", data)
	}
}

func testUnidirectionalChannel() {
	// 定义双向channel
	cha := make(chan int)
	// 写数据
	go writeData(cha)
	// 读数据
	go readData(cha)

	time.Sleep(1 * time.Second)
}

func TestChannel() {
	fmt.Println("--------no buffer channel---------")
	testNobufferChannel()

	fmt.Println("--------buffer channel---------")
	bufferCap := 3
	testBufferChannel(bufferCap)

	fmt.Println("--------closed channel---------")
	testCloseChannel()

	fmt.Println("--------unidirectional channel---------")
	testUnidirectionalChannel()
}
