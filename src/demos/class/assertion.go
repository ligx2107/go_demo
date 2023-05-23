package cla

import "fmt"

/*
	  断言：
		1. 万能类型：interface{}可以作为任何类型数据容器
		2. 万能类型可通过断言语法，判断引用的底层真实数据类型
*/
func assertionFunc(arg interface{}) {
	fmt.Println("assertionFunc called...")
	// 通过断言，判断真实数据类型, value -> 万能类型变量的值， ok -> 断言是否为真
	value, ok := arg.(string)
	if ok {
		fmt.Printf("arg's value %s\n", value)
		fmt.Printf("arg's type %T\n", arg)
	} else {
		fmt.Println("arg is other type")
	}
}

func TestAssertion() {
	arg1 := 1
	assertionFunc(arg1)

	arg2 := "hello"
	assertionFunc(arg2)
}
