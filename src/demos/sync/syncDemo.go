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
	sync.Once：并发场景下，保证特定代码只被执行一次
		1. func (o *Once) Do(f func())
		2. 只有在当前Once实例第一次调用Do方法时，才会真正执行f函数
		3. 在对f函数的调用返回之前，不会返回对Do的调用，所以在f函数中再次调用Do方法，会产生死锁
		4. 如果f函数抛出了panic，此时Do会认为f函数已经返回
	sync.Mutex: 互斥锁
	sync.RWMutex: 读写锁, 在读多写少的场景下，效率高于互斥锁
		1. 互斥场景：读写，写读，写写
	sync.Map: 并发安全的Map，开箱即用，不需要通过make函数进行初始化
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

func testOnce() {
	f := func(arg int) {
		fmt.Println("anonymous function execute, arg -->", arg)
	}

	var wg sync.WaitGroup
	var executeOnce sync.Once
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			executeOnce.Do(func() {
				f(2)
			}) // 调用sync.Once.Do方法，在并发场景下保证f函数只执行一次
		}()
	}

	wg.Wait()
}

func testMutex() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 5000; j++ {
				lock.Lock()
				r += 1 //在不加锁的情况下，本行代码是非并发安全的
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("r value: ", r)
}

func TestSync() {
	testWaitGroup()
	testOnce()
	testMutex()
}
