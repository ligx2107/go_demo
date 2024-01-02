package gr

import "runtime"

/**
 *  Gosched 暂停，释放线程去执行其他任务，当前任务被放回队列等待下次被执行
 *  Goexit 立即终止当前任务，运行时确保所有已注册defer延迟调用被执行，不会影响其他并发任务，不会引发panic
 *         无论身处哪一层，Goexit都能立即终止整个调用栈，与return仅退出当前函数不同，os.Exit可终止进程，但不会执行延迟调用defer
 */

func TestGosched() {
	runtime.GOMAXPROCS(1)       // 单线程执行
	exit := make(chan struct{}) // 声明一个空结构体类型的channel

	go func() {
		defer close(exit)

		go func() {
			println("b")
		}()

		for i := 0; i < 4; i++ {
			println("a:", i)
			if i == 2 {
				runtime.Gosched()
			}
		}
	}()

	<-exit // 通过channel等待子协程结束
}

func TestGoexit() {

	runtime.GOMAXPROCS(1)       // 单线程执行
	exit := make(chan struct{}) // 声明一个空结构体类型的channel

	go func() {
		defer close(exit)  // 会执行
		defer println("a") // 会执行

		func() {
			defer func() {
				println("b", recover() == nil)
			}()

			func() {
				println("c")
				runtime.Goexit()  // 立即终止整个调用栈
				println("c done") // 不会执行
			}()

			println("b done") // 不会执行
		}()

		println("a done") // 不会执行
	}()

	<-exit // 通过channel等待子协程结束
}
