package deferDemo

import "fmt"

/**
  defer: Golang提供的一种用于注册延迟调用的机制，让函数或语句可以在当前函数执行完毕后(包括通过return正常结束)执行
	1. 多个defer执行顺序: defer语句采用压栈方式，LIFO，后声明的先被调用
	2. defer与return的执行顺序：return先于defer执行
*/

func test1() {
	fmt.Println("test1 called...")
}

func test2() {
	fmt.Println("test2 called...")
}

func deferFun() {
	fmt.Println("deferFun called...")
}

func returnFun() int {
	fmt.Println("returnFun called...")
	return 1
}

func DeferFunc() {
	// 声明两个defer，采用压栈方式，后入先出，test2先于test1被执行
	defer test1()
	defer test2()
}

func DeferReturnFunc() int {
	// return先于defer被执行
	defer deferFun()
	return returnFun()
}

type Person struct {
	name string
}

func (p *Person) printName() {
	fmt.Println("person name: ", p.name)
}

func DeferClosure() {
	persons := []Person{Person{"zhangsan"}, Person{"lisi"}, Person{"wangwu"}}
	for _, p := range persons {
		defer p.printName()
	}

	for i := 0; i < len(persons); i++ {
		defer persons[i].printName()
	}
}
