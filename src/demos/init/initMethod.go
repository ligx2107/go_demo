package init

import "time"

/**
 *	1. 包内每个源码文件都可以包含一到多个init函数，但编译器不保证执行顺序
 *	2. 包内所有init函数都由编译器自动生成的一个包装函数进行调用，因此可保证在单一线程上执行，且仅执行一次
 *	3. 直到全部init函数执行完毕后，才会执行main.main函数
 *	4. init函数不能被其他函数调用
 */
func init() {
	println("init method 1")
}

func init() {
	done := make(chan int)
	go func() {
		defer close(done)

		println("init method 2")
		time.Sleep(2 * time.Second)
	}()

	<-done
}
