package atomicdemo

import (
	"fmt"
	"sync/atomic"
)

func test() {
	var count int32 = 10
	// 设置值
	atomic.StoreInt32(&count, 20)
	fmt.Println("设置后的count值为：", atomic.LoadInt32(&count)) // 20

	// 相加
	atomic.AddInt32(&count, 10)
	fmt.Println("相加后的count值为：", atomic.LoadInt32(&count)) // 30

	// 相减
	atomic.AddInt32(&count, -10)
	fmt.Println("相减后的count值为：", atomic.LoadInt32(&count)) // 20

	// 交换变量的值并返回原有的值
	fmt.Println("交换后的原始值：", atomic.SwapInt32(&count, 40)) // 20
	fmt.Println("交换后的值：", atomic.LoadInt32(&count))       // 40

	// 先比较后交换
	fmt.Println("是否交换成功：", atomic.CompareAndSwapInt32(&count, 41, 50)) // true
	fmt.Println("是否交换成功：", atomic.CompareAndSwapInt32(&count, 40, 50)) // true

}

func TestAtomic() {
	test()
}
