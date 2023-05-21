package fun

import "fmt"

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
