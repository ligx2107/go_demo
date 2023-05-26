package closure

import "fmt"

// 定义test函数，返回值为两个函数类型变量
func test(base int) (func(int) int, func(int) int) {
	add := func(p1 int) int {
		base += p1
		return base
	}
	sub := func(p1 int) int {
		return base - p1
	}

	return add, sub
}

func TestClosure() {
	// add，sub两个闭包，同时拥有test函数作用域范围内的base变量引用
	add, sub := test(10)
	fmt.Println(add(1)) // 11
	fmt.Println(sub(2)) // 9
}
