package syncd

import (
	"fmt"
	"sync"
	"time"
)

var r int64

/*
	sync.WaitGroup通过内部维护的计数器，实现并发任务的同步
		1. Add: 计数器增加指定值
		2. Done: 计数器减1
		3. Wait: 阻塞直到计数器为0
		4. 协程间传递时需要以指针或者闭包的方式引用WaitGroup对象，否则会造成死锁
	sync.Cond：
		1. 设置一组协程根据条件阻塞，可以根据不同的条件阻塞
		2. 可以根据条件唤醒相对应的协程
		3. 适用于一发多收的场景
	sync.Once：并发场景下，保证特定代码只被执行一次
		1. func (o *Once) Do(f func())
		2. 只有在当前Once实例第一次调用Do方法时，才会真正执行f函数
		3. 在对f函数的调用返回之前，不会返回对Do的调用，所以在f函数中再次调用Do方法，会产生死锁
		4. 如果f函数抛出了panic，此时Do会认为f函数已经返回
	sync.Mutex: 互斥锁
	sync.RWMutex: 读写锁, 在读多写少的场景下，效率高于互斥锁
		1. 互斥场景：读写，写读，写写
	sync.Map: 并发安全的Map，开箱即用，不需要通过make函数进行初始化
	sync.Pool: 创建一个临时对象池，缓存一组对象用于重复利用，用来减少内存分配和降低GC压力
		1. 用于缓存一些创建成本较高且使用比较频繁的对象
		2. Pool的长度默认为机器CPU的线程数
		3. 存储在Pool中的对象随时都有可能被回收
		4. 创建成本低的对象，不建议使用对象池
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

func testCond() {
	list := make([]int, 0)
	cond := sync.NewCond(&sync.Mutex{})

	// 启用协程读取list中数据
	go readList(&list, cond)
	go readList(&list, cond)
	go readList(&list, cond)

	// 模拟
	time.Sleep(time.Second * 2)

	// 主协程写入数据
	initList(&list, cond)
}

func initList(list *[]int, cond *sync.Cond) {
	// 主叫方，可以持有锁，也可以不持锁
	//cond.L.Lock()
	//defer cond.L.Unlock()

	for i := 1; i <= 10; i++ {
		*list = append(*list, i)
	}

	// 唤醒所有条件等待的协程
	cond.Broadcast()

	// 唤醒其中一个条件等待协程
	//cond.Signal()
}

func readList(list *[]int, cond *sync.Cond) {
	// 被叫方，必须持有锁
	cond.L.Lock()
	defer cond.L.Unlock()
	for len(*list) == 0 {
		fmt.Println("read list waiting...")
		// 等待被唤醒
		cond.Wait()
	}

	fmt.Println("list数据为: ", *list)
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
	testCond()
}
