package err

import (
	"fmt"
)

/*
	panic:
		1. panic错误会中断程序执行并按调用栈传递直到整个程序退出
		2. panic所在的函数内如果存在defer语句，则defer可正常执行
	recover:
		1. 用来控制程序的panic行为，捕获panic传递的error，恢复正常代码的执行
		2. recover只有在panic之前定义的defer函数中直接调用才会捕获panic错误
		3. recover处理异常后，程序会跳转到defer之后的点继续执行
		4. defer中引发的错误，可被后续的defer调用捕获，按defer调用栈的执行顺序，后一个执行的defer可捕获前一个defer中出现的panic
*/

func panicDemo(num1, num2 int) interface{} {
	var r interface{}
	defer func() {
		fmt.Println("defer called...")
		if error := recover(); error != nil {
			fmt.Println("error message: ", error)
		}
	}()
	r = num1 / num2
	fmt.Println("r -->", r)
	return r
}

func multipleDefer() {
	defer func() {
		fmt.Println("error info:", recover())
	}()

	defer func() {
		fmt.Println("catch error info: ", recover())
		panic("defer one")
	}()

	defer func() {
		panic("defer two")
	}()

	panic("panic error")
}

func TestPanic() {
	fmt.Println(panicDemo(5, 0))
	fmt.Println("after panic")

	multipleDefer()
}
