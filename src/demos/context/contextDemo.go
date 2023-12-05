package contextdemo

import (
	"context"
	"fmt"
	"time"
)

/*
	context:
		1. 用于进程间信息和信号传递
		2. 用于服务间信息和信号传递
		3. 用于父子协程间取消信号传递
		4. 用于设置请求超时等
*/

func TestContext() {
	test1()
	test2()
}

// 父子协程间的信号传递
func test2() {
	// Background方法，返回一个非nil的空context
	ctx := context.Background()

	// WithValue方法，返回参数context的副本，同时带有所设置的key：value
	ctx = context.WithValue(ctx, "desc", "test2")

	// 设置超时时间
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)

	// 延时调用cancel方法
	defer cancel()

	// 定义测试数据
	data := [][]int{{2, 2}, {3, 3}}

	// 定义chan
	ch := make(chan []int)

	// 启动协程，调用计算方法
	go calculate(ctx, ch)

	// 向chan中写入数据
	for i := 0; i < len(data); i++ {
		ch <- data[i]
	}
}

// 定义计算方法
func calculate(ctx context.Context, ch <-chan []int) {
	// 监听chan
	for {
		select {
		case data := <-ch:
			ctx := context.WithValue(ctx, "desc", "calculate")

			// 求和
			sumCh := make(chan []int)
			go sumContext(ctx, sumCh)
			sumCh <- data

			// 求积
			multCh := make(chan []int)
			go multContext(ctx, multCh)
			multCh <- data

		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("协程calculate退出，context desc：%s，错误消息：%s\n", desc, ctx.Err())
			return
		}
	}
}

func sumContext(ctx context.Context, ch <-chan []int) {
	for {
		select {
		case data := <-ch:
			a1, a2 := data[0], data[1]
			res := sum(a1, a2)
			fmt.Printf("%d + %d = %d\n", a1, a2, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("协程sumContext退出，context desc: %s, 错误信息：%s\n", desc, ctx.Err())
			return
		}
	}
}

func multContext(ctx context.Context, ch <-chan []int) {
	for {
		select {
		case data := <-ch:
			a1, a2 := data[0], data[1]
			res := mult(a1, a2)
			fmt.Printf("%d * %d = %d\n", a1, a2, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("协程multContext退出，context desc: %s, 错误信息：%s\n", desc, ctx.Err())
			return
		}
	}
}

func sum(a1, a2 int) int {
	return a1 + a2
}

func mult(a1, a2 int) int {
	return a1 * a2
}

// 通过关闭chan的方式，通知协程退出
func test1() {
	done := make(chan struct{})
	// 启动协程
	go f1(done)
	go f1(done)

	// 主协程等待
	time.Sleep(time.Second * 3)
	// 关闭chan
	close(done)
}

func f1(done chan struct{}) {
	// 监听chan
	for {
		select {
		case <-done:
			fmt.Println("协程退出")
			return
		}
	}
}
