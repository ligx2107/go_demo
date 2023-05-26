package fun

import "fmt"

/**
  golang函数特点：
	1. 支持变参，支持多返回值，支持命名返回参数，支持匿名函数和闭包，可作为一种类型赋值给变量
	2. 不支持重载(同一个包内不能有同名函数)，不支持默认参数
  方法命名：
	1. 首字母大写，为包向外暴露的方法，供其他包程序调用
	2. 首字母小写，为包内私有方法，只供同一个包内程序调用
*/

// 返回多个匿名返回值
func multipleAnonymousReturn(p1 string, p2 int) (int, string) {
	fmt.Println("MultipleAnonymousReturn called...")
	return 100, "hello " + p1
}

// 返回多个具名返回值
func multipleNamedReturn(p1 string, p2 int) (r1 int, r2 string) {
	fmt.Println("MultipleNamedReturn called...")

	//为具名返回值变量赋值
	r1 = p2
	r2 = "hello " + p1
	return
}

// 定义函数类型变量
type printFunc func(format string, n1, n2 int)

// 定义接收函数类型参数的函数
func testFunctionType(fn printFunc, format string, n1, n2 int) {
	fn(format, n1, n2)
}

func TestFunction() {
	r1, r2 := multipleAnonymousReturn("world", 100)
	fmt.Println("r1 =", r1, "r2 =", r2)
	r3, r4 := multipleNamedReturn("Golang", 200)
	fmt.Println("r3 =", r3, "r4 =", r4)

	// 匿名函数作函数类型参数
	testFunctionType(func(format string, n1, n2 int) {
		fmt.Printf(format, n1, n2)
	}, "format result --> n1: %d, n2: %d\n", 2, 3)

	// 变参函数
	s1 := []int{1, 2, 3}
	func(arg ...int) {
		fmt.Println("arg length: ", len(arg))
		for i, v := range arg {
			fmt.Printf("arg[%d], value: %d\n", i, v)
		}
	}(s1...) //用slice对象作变参时，必须要展开
}
