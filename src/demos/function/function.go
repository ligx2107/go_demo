package fun

import "fmt"

/**
  方法命名：
	1. 首字母大写，为包向外暴露的方法，供其他包程序调用
	2. 首字母小写，为包内私有方法，只供同一个包内程序调用
*/

// 返回多个匿名返回值
func MultipleAnonymousReturn(p1 string, p2 int) (int, string) {
	fmt.Println("MultipleAnonymousReturn called...")
	return 100, "hello " + p1
}

// 返回多个具名返回值
func MultipleNamedReturn(p1 string, p2 int) (r1 int, r2 string) {
	fmt.Println("MultipleNamedReturn called...")

	//为具名返回值变量赋值
	r1 = p2
	r2 = "hello " + p1
	return
}
