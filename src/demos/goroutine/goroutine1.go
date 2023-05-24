package gr

import (
	"fmt"
	"runtime"
	"time"
)

/*
	goroutine:
		1. 创建goroutine，通过go关键字执行一个函数(普通函数，匿名函数)，一个goroutine必定与一个函数对应
		2. goroutine调度:
			2.1. GPM：GPM是Go语言运行时(runtime)层面的实现，是Go语言自己实现的一套调度系统
			2.2. G(goroutine)，除了存放goroutine本身的信息外，还存放与P的绑定关系
			2.3. P(processor), 管理着一组goroutine，同时存储当前goroutine运行的上下文环境信息(函数指针，堆栈地址等)，会对自己所管理的goroutine队列做一些调度
			2.4. M(machine), Go运行时(runtime)对操作系统内核线程的虚拟，M与内核线程是一对一映射关系
			2.5. P与M一般也是一对一的，P管理着一组G挂载在M上运行
			2.6. P的个数可以通过runtime.GOMAXPROCS设定
		3. goroutine的调度是在用户态下完成的，不涉及用户态和内核态的切换
*/

func TestGoroutine() {
	// 用go创建一个goroutine
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 退出当前goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
