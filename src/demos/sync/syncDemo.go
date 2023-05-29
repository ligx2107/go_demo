package syncd

import (
	"fmt"
	"sync"
)

var r int64

/*
	sync.WaitGroup通过内部维护的计数器，实现并发任务的同步
		1. Add: 计数器增加指定值
		2. Done: 计数器减1
		3. Wait: 阻塞直到计数器为0
*/

func testWaitGroup() {
	var wg sync.WaitGroup //使用waitGroup实现goroutine间的同步
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine execute....")
	}()
	wg.Wait()
	fmt.Println("main goroutine execute...")
}

func TestSync() {
	testWaitGroup()
}
